package system

import (
	"errors"
	"sync"

	"github.com/iancoleman/orderedmap"
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

func (e *Events) CallEvents(engine interface{}, name string) error {
	res, err := e.GetEvents(name)
	if err != nil {
		return nil
	}
	for _, event := range res {
		err := event(engine)
		if err != nil {
			return errors.New("Event error: " + err.Error())
		}
	}
	return nil
}

func NewEvents() *Events {
	return &Events{
		Scope:  make(ScopeType),
		events: *orderedmap.New(),
	}
}
