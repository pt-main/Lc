package core

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/iancoleman/orderedmap"
	"github.com/pt-main/lc/public"
)

type EventsInterface interface {
	GetEvents(name string) ([]EventType, error)
	NewEvent(name string, event EventType)
	NewEventBefore(name string, event EventType) error
	CallEvents(input *EventInput, name string, canWorkWithoutHandler bool) error
	Scope() ScopeType
}

// Events manages an ordered collection of event handlers. Each event has a
// name (string) and a list of EventType functions. The CallEvents method
// invokes all handlers of an event in registration order. Events can also
// automatically wrap calls with start/end events for logging.
type Events struct {
	scope   ScopeType
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

func (e *Events) NewEventBefore(name string, event EventType) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	list, err := e.GetEvents(name)
	if err != nil {
		return fmt.Errorf("Can't put new event before '%s': %v", name, err)
	}
	e.events.Set(name, append([]EventType{event}, list...))
	return nil
}

func (e *Events) callEvents(input *EventInput, name string, canWorkWithoutHandler bool) error {
	res, err := e.GetEvents(name)
	if err != nil {
		if canWorkWithoutHandler {
			return nil
		} else {
			return errors.New("Event '" + name + "' not found.")
		}
	}
	for _, event := range res {
		err := event(e, input)
		if err != nil {
			return errors.New("Event error: " + err.Error())
		}
	}
	return nil
}

func (e *Events) CallEvents(input *EventInput, name string,
	canWorkWithoutHandler bool) error {
	e.scope[public.EventsScopeCallName] = name
	var err error
	err = e.callEvents(nil, public.CallEventsStartEvent, true)
	if err != nil {
		return err
	}
	err = e.callEvents(input, name, canWorkWithoutHandler)
	e.scope[public.EventsScopeCallError] = err
	err1 := e.callEvents(nil, public.CallEventsEndEvent, true)
	if err1 != nil {
		return err1
	}
	if err != nil {
		return err
	}
	return nil

}

func (e *Events) Scope() ScopeType {
	return e.scope
}

// NewEvents creates an empty Events instance with an ordered map.
// The Scope map is initially empty but can be used to pass data between
// event handlers.
func NewEvents(context context.Context) *Events {
	return &Events{
		scope:   make(ScopeType),
		events:  *orderedmap.New(),
		Context: context,
	}
}
