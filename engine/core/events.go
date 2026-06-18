package core

import (
	"context"
	"errors"
	"sync"

	"github.com/iancoleman/orderedmap"
)

const (
	CallEventsStartEvent = "->call(Events.CallEvents)"
	CallEventsEndEvent   = "call(Events.CallEvents)->"
)

// Events manages an ordered collection of event handlers. Each event has a
// name (string) and a list of EventType functions. The CallEvents method
// invokes all handlers of an event in registration order. Events can also
// automatically wrap calls with start/end events for logging.
type Events struct {
	Scope   ScopeType
	Context context.Context
	mu      sync.RWMutex
	events  orderedmap.OrderedMap
}

func (e *Events) GetEvents(name string) ([]EventType, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	val, ok := e.events.Get(name)
	if !ok {
		return nil, errors.New("Can't find event '" + name + "'")
	}
	list, _ := val.([]EventType)
	return list, nil
}

func (e *Events) NewEvent(name string, event EventType) {
	e.mu.Lock()
	defer e.mu.Unlock()

	val, ok := e.events.Get(name)
	if !ok {
		e.events.Set(name, []EventType{event})
		return
	}
	list, ok := val.([]EventType)
	if !ok {
		e.events.Set(name, []EventType{event})
		return
	}
	e.events.Set(name, append(list, event))
}

func (e *Events) callEvents(input interface{}, name string,
	canWorkWithoutHandler bool) error {
	res, err := e.GetEvents(name)
	if err != nil {
		if canWorkWithoutHandler {
			return nil
		} else {
			return errors.New("Event '" + name + "' not found.")
		}
	}
	for _, event := range res {
		err := event(input, e)
		if err != nil {
			return errors.New("Event error: " + err.Error())
		}
	}
	return nil
}

const (
	EventsScopeCallName  = "CALL_NAME string"
	EventsScopeCallError = "CALL_ERROR error"
)

func (e *Events) CallEvents(input interface{}, name string,
	canWorkWithoutHandler bool) error {
	e.Scope[EventsScopeCallName] = name
	var err error
	err = e.callEvents(name, CallEventsStartEvent, true)
	if err != nil {
		return err
	}
	err = e.callEvents(input, name, canWorkWithoutHandler)
	e.Scope[EventsScopeCallError] = err
	err1 := e.callEvents(name, CallEventsEndEvent, true)
	if err1 != nil {
		return err1
	}
	if err != nil {
		return err
	}
	return nil

}

// NewEvents creates an empty Events instance with an ordered map.
// The Scope map is initially empty but can be used to pass data between
// event handlers.
func NewEvents(context context.Context) *Events {
	return &Events{
		Scope:   make(ScopeType),
		events:  *orderedmap.New(),
		Context: context,
	}
}
