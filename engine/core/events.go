package core

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/pt-main/lc/public"
)

type EventsInterface interface {
	GetEvents(name string) ([]EventType, error)
	GetCoreEventIdx(name string) (int, error)
	SetEvents(name string, events []EventType, idx int)
	NewEvent(name string, event EventType)
	NewEventBefore(name string, event EventType) error
	CallEvents(input *EventInput, name string, canWorkWithoutHandler bool) error
	Scope() ScopeType
	CoreEvents() map[string]int
}

// Events manages an ordered collection of event handlers. Each event has a
// name (string) and a list of EventType functions. The CallEvents method
// invokes all handlers of an event in registration order. Events can also
// automatically wrap calls with start/end events for logging.
type Events struct {
	scope      ScopeType
	Context    context.Context
	mu         sync.RWMutex
	coreEvents map[string]int
	events     map[string][]EventType
}

func (e *Events) GetEvents(name string) ([]EventType, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	val, ok := e.events[name]
	if !ok {
		return nil, errors.New("Can't find event '" + name + "'")
	}
	return val, nil
}

func (e *Events) SetEvents(name string, events []EventType, coreEvent int) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if events == nil {
		events = make([]EventType, 0)
	}
	e.events[name] = events
	e.coreEvents[name] = coreEvent
}
func (e *Events) GetCoreEventIdx(name string) (int, error) {
	ce, ok := e.coreEvents[name]
	if !ok {
		return -1, fmt.Errorf("Can't find event: %v", name)
	}
	return ce, nil
}

func (e *Events) CoreEvents() map[string]int {
	return e.coreEvents
}

func (e *Events) NewEvent(name string, event EventType) {
	e.mu.Lock()
	defer e.mu.Unlock()

	val, ok := e.events[name]
	if !ok {
		e.events[name] = []EventType{event}
		e.coreEvents[name] = 0
		return
	}
	e.events[name] = append(val, event)
}

func (e *Events) NewEventBefore(name string, event EventType) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	list, err := e.GetEvents(name)
	if err != nil {
		return fmt.Errorf("Can't put new event before '%s': %v", name, err)
	}
	e.events[name] = append([]EventType{event}, list...)
	e.coreEvents[name] += 1
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
		scope:      make(ScopeType),
		events:     make(map[string][]EventType),
		coreEvents: make(map[string]int),
		Context:    context,
	}
}

type EventsTools struct {
	Events EventsInterface
}

func (et *EventsTools) ChangeCoreEvent(name string, event EventType) error {
	e := et.Events
	idx, err := e.GetCoreEventIdx(name)
	if err != nil {
		return err
	}
	events, err := e.GetEvents(name)
	if err != nil {
		return err
	}
	if idx <= 0 {
		e.SetEvents(name, []EventType{event}, 0)
		return nil
	}
	pre := events[:idx]
	post := events[idx+1:]
	done := append(append(pre, event), post...)
	e.SetEvents(name, done, idx)
	return nil
}

func (et *EventsTools) GetCoreEvent(name string) (EventType, error) {
	var etn EventType
	e := et.Events
	idx, err := e.GetCoreEventIdx(name)
	if err != nil {
		return etn, err
	}
	ev, err := e.GetEvents(name)
	if err != nil {
		return etn, err
	}
	if idx < 0 || idx > (len(ev)-1) {
		return etn, fmt.Errorf("Invalid core event idx")
	}
	return ev[idx], nil
}
