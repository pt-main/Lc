package core

import (
	"context"
	"sync"

	"github.com/pt-main/lc/v2/public"
	"github.com/pt-main/lc/v2/public/errors"
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
// name and a list of EventType functions, and CallEvents invokes the
// handlers of an event in registration order.
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
		return nil, Err(errors.EventsEventIsNotFound, "%s", name)
	}
	return append(make([]EventType, 0, len(val)), val...), nil
}

func (e *Events) SetEvents(name string, events []EventType, coreEvent int) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if events == nil {
		events = []EventType{}
	}
	// A copy is stored, otherwise a caller keeping the slice could mutate the
	// handler list without holding e.mu.
	e.events[name] = append(make([]EventType, 0, len(events)), events...)
	e.coreEvents[name] = coreEvent
}

// Err errors.EventsEventIsNotFound, Msg=name
func (e *Events) GetCoreEventIdx(name string) (int, ErrorInterface) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	ce, ok := e.coreEvents[name]
	if !ok {
		return -1, Err(errors.EventsEventIsNotFound, "%s", name)
	}
	return ce, nil
}

func (e *Events) CoreEvents() map[string]int {
	e.mu.RLock()
	defer e.mu.RUnlock()
	res := make(map[string]int, len(e.coreEvents))
	for k, v := range e.coreEvents {
		res[k] = v
	}
	return res
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
		}
		return Wrap(errors.EventsEventError, err, "Can't find event")
	}
	for _, event := range res {
		if err := event(e, input); err != nil {
			return Wrap(errors.EventsEventError, err, "Event handler failed")
		}
	}
	return nil
}

// Err errors.EventsEventError (from 'callEvents')
func (e *Events) CallEvents(input *EventInput, name string, canWorkWithoutHandler bool) ErrorInterface {
	// e.mu guards the maps in this struct, so the two bookkeeping writes go
	// through it; without it concurrent calls raced on the shared scope map.
	ScopeSetSynced(e.scope, public.EventsScopeCallName, name)
	var err ErrorInterface
	if e.isDebug() {
		if err = e.callEvents(nil, public.CallEventsStartEvent, true); err != nil {
			return err
		}
	}
	err = e.callEvents(input, name, canWorkWithoutHandler)
	ScopeSetSynced(e.scope, public.EventsScopeCallError, err)
	if e.isDebug() {
		// A failure of the debug end handler must not hide the real error, so
		// it is only reported when the call itself succeeded.
		if err1 := e.callEvents(nil, public.CallEventsEndEvent, true); err1 != nil && err == nil {
			return err1
		}
	}
	return err
}

func (e *Events) isDebug() bool {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.debug
}

func (e *Events) ReplaceEvent(name string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	delete(e.events, name)
	delete(e.coreEvents, name)
}

// Err errors.EventsSystemError
func (e *Events) SetProperty(name string, value interface{}) ErrorInterface {
	switch name {
	case "debug":
		flag, ok := value.(bool)
		if !ok {
			return Err(errors.EventsSystemError, "Invalid property value (must be bool): %v", value)
		}
		e.mu.Lock()
		defer e.mu.Unlock()
		e.debug = flag
		return nil
	default:
		return Err(errors.EventsSystemError, "Invalid property name: %v", name)
	}
}

func (e *Events) Scope() ScopeType {
	return e.scope
}

// NewEvents creates an empty Events instance. The Scope map starts empty and
// is shared by the event handlers.
func NewEvents(ctx context.Context) *Events {
	return &Events{
		scope:      make(ScopeType),
		events:     make(map[string][]EventType),
		coreEvents: make(map[string]int),
		Context:    ctx,
	}
}

// EventsTools works with the core event of an event, the first handler
// registered under its name.
type EventsTools struct {
	Events EventsInterface
}

// ChangeCoreEvent replaces the core event of name, leaving every other
// handler of that event in place.
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
	done := append(append(events[:idx:idx], event), events[idx+1:]...)
	e.SetEvents(name, done, idx)
	return nil
}

// GetCoreEvent returns the core event of name.
func (et *EventsTools) GetCoreEvent(name string) (EventType, ErrorInterface) {
	e := et.Events
	idx, err := e.GetCoreEventIdx(name)
	if err != nil {
		return nil, err
	}
	ev, err := e.GetEvents(name)
	if err != nil {
		return nil, err
	}
	if idx < 0 || idx > len(ev)-1 {
		return nil, Err(errors.EventsSystemError, "Invalid core event idx")
	}
	return ev[idx], nil
}
