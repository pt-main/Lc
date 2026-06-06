package core

import (
	"errors"
	"sync"

	"github.com/iancoleman/orderedmap"
)

const (
	CallEventsEvent = "call(Events.CallEvents)"
)

type Events struct {
	Scope  ScopeType
	mu     sync.RWMutex
	events orderedmap.OrderedMap
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
		err := event(input)
		if err != nil {
			return errors.New("Event error: " + err.Error())
		}
	}
	return nil
}

func (e *Events) CallEvents(input interface{}, name string,
	canWorkWithoutHandler bool) error {
	e.Scope["event_input"] = input
	var err error
	err = e.callEvents(name, CallEventsEvent, true)
	if err != nil {
		return err
	}
	err = e.callEvents(input, name, canWorkWithoutHandler)
	if err != nil {
		return err
	}
	return nil

}

func NewEvents() *Events {
	return &Events{
		Scope:  make(ScopeType),
		events: *orderedmap.New(),
	}
}
