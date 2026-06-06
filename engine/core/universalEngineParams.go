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
	logS := func(_name interface{}) error {
		name, ok := _name.(string)
		if !ok {
			return fmt.Errorf("Invalid event input: bad interface")
		}
		logger.PrintLog("debug", "Start call '"+name+"' event")
		return nil
	}
	logE := func(_name interface{}) error {
		name, ok1 := _name.(string)
		err, ok2 := _name.(error)
		if !ok1 || !ok2 {
			return fmt.Errorf("Invalid event input: bad interface")
		}
		text := "End call '" + name + "' event"
		if err != nil {
			text += " with error: " + err.Error()
		}
		logger.PrintLog("debug", text)
		return nil
	}
	event.NewEvent(CallEventsStartEvent, logS)
	event.NewEvent(CallEventsEndEvent, logE)
	return &UniversalEngineParams{
		Generator: generator,
		Event:     event,
		Scope:     scope,
		Logger:    logger,
	}, nil
}
