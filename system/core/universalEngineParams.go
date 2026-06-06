package core

import "fmt"

type UniversalEngineParams struct {
	Generator *Generator
	Event     *Events
	Scope     ScopeType
	Logger    *Logger
}

func NewUniversalEngineParams(
	generator *Generator,
	event *Events,
	scope ScopeType,
	logger *Logger,
) (*UniversalEngineParams, error) {
	log := func(_name interface{}) error {
		name, ok := _name.(string)
		if !ok {
			return fmt.Errorf("Invalid event input: bad interface")
		}
		logger.PrintLog("debug", "Call "+name+"event")
		return nil
	}
	event.NewEvent(CallEventsEvent, log)
	return &UniversalEngineParams{
		Generator: generator,
		Event:     event,
		Scope:     scope,
		Logger:    logger,
	}, nil
}
