package core

import (
	"context"
	"sync"

	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/public/errors"
)

type EventsInterface interface {
	GetEvents(name string) ([]EventType, ErrorInterface)
	GetCoreEventIdx(name string) (int, ErrorInterface)
	SetEvents(name string, events []EventType, idx int)
	NewEvent(name string, event EventType)
	NewEventBefore(name string, event EventType) ErrorInterface
	CallEvents(input *EventInput, name string, canWorkWithoutHandler bool) ErrorInterface
	Scope() ScopeType
	CoreEvents() map[string]int
	ReplaceEvent(name string)
	SetProperty(name string, value interface{}) ErrorInterface
}

// Events manages an ordered collection of event handlers. Each event has a
// name (string) and a list of EventType functions. The CallEvents method
// invokes all handlers of an event in registration order. Events can also
// automatically wrap calls with start/end events for logging.
type Events struct {
	scope      ScopeType
	Context    context.Context
	mu         sync.RWMutex
	debug      bool
	coreEvents map[string]int
	events     map[string][]EventType
}

// Err errors.EventsEventIsNotFound, Msg=name
func (e *Events) GetEvents(name string) ([]EventType, ErrorInterface) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	val, ok := e.events[name]
	if !ok {
		return nil, Err(errors.EventsEventIsNotFound, name)
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

// Err errors.EventsEventIsNotFound, Msg=name
func (e *Events) GetCoreEventIdx(name string) (int, ErrorInterface) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	ce, ok := e.coreEvents[name]
	if !ok {
		return -1, Err(errors.EventsEventIsNotFound, name)
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

// Err errors.EventsSystemError.
// With meta: EMK(0, "string") - event name, EMK(1, "error") (from 'GetEvents')
func (e *Events) NewEventBefore(name string, event EventType) ErrorInterface {
	list, err := e.GetEvents(name)
	if err != nil {
		return Wrap(errors.EventsSystemError, err, "Can't put new event before '%s'", name).
			WithMeta(EMK(0, "string"), name)
	}
	e.mu.Lock()
	defer e.mu.Unlock()
	e.events[name] = append([]EventType{event}, list...)
	e.coreEvents[name] += 1
	return nil
}

// Err errors.EventsEventError.
// Cause: errors.EventsEventIsNotFound (from 'GetEvents') or event error, msg = event ErrorInterface text
func (e *Events) callEvents(input *EventInput, name string, canWorkWithoutHandler bool) ErrorInterface {
	res, err := e.GetEvents(name)
	if err != nil {
		if canWorkWithoutHandler {
			return nil
		} else {
			return Wrap(errors.EventsEventError, err, "Can't found event")
		}
	}
	for _, event := range res {
		err := event(e, input)
		if err != nil {
			return Wrap(errors.EventsEventError, err, "Event handler failed")
		}
	}
	return nil
}

// Err errors.EventsEventError (from 'callEvents')
func (e *Events) CallEvents(input *EventInput, name string,
	canWorkWithoutHandler bool) ErrorInterface {
	e.scope[public.EventsScopeCallName] = name
	var err ErrorInterface
	if e.debug {
		err = e.callEvents(nil, public.CallEventsStartEvent, true)
		if err != nil {
			return err
		}
	}
	err = e.callEvents(input, name, canWorkWithoutHandler)
	e.scope[public.EventsScopeCallError] = err
	if e.debug {
		err1 := e.callEvents(nil, public.CallEventsEndEvent, true)
		if err1 != nil {
			return err1
		}
	}
	if err != nil {
		return err
	}
	return nil

}

func (e *Events) ReplaceEvent(name string) {
	delete(e.events, name)
	delete(e.coreEvents, name)
}

// Err errors.EventsSystemError
func (e *Events) SetProperty(name string, value interface{}) ErrorInterface {
	switch name {
	case "debug":
		var ok bool
		e.debug, ok = value.(bool)
		if !ok {
			return Err(errors.EventsSystemError, "Invalid property value (must be bool): %v", value)
		}
		return nil
	default:
		return Err(errors.EventsSystemError, "Invalid property name: %v", name)
	}
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

func (et *EventsTools) ChangeCoreEvent(name string, event EventType) ErrorInterface {
	e := et.Events
	idx, err := e.GetCoreEventIdx(name)
	if err != nil {
		return err
	}
	events, err := e.GetEvents(name)
	if err != nil {
		return err
	}
	if idx < 0 {
		return Err(errors.EventsSystemError, "Can't change core event: %v", name)
	}
	pre := events[:idx]
	post := events[idx+1:]
	done := append(append(pre, event), post...)
	e.SetEvents(name, done, idx)
	return nil
}

func (et *EventsTools) GetCoreEvent(name string) (EventType, ErrorInterface) {
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
		return etn, Err(errors.EventsSystemError, "Invalid core event idx")
	}
	return ev[idx], nil
}
