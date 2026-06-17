package core

import (
	"context"
	"fmt"
)

// UniversalEngineParams is a container shared by both StringEngine and ByteEngine.
// It holds the Generator (for output accumulation), Events (for hooks),
// Scope (for passing arbitrary data between stages), and Logger (for diagnostics).
// This struct is embedded (not composed by pointer) in the engine types,
// promoting its fields and methods to the engine itself.
type UniversalEngineParams struct {
	// Generator *Generator - controls code/output generation across pipeline points.
	Generator *Generator
	// Event *Events - allows hooking into parsing and command dispatch.
	Event *Events
	// Scope ScopeType - a map[string]interface{} that can be used to share
	//   variables between events, parsers, and command handlers.
	Scope ScopeType
	// Logger *Logger - if set, logs internal steps (event calls, errors).
	Logger *Logger
	// If true - enable colored messages in enine (like errors, etc.).

	Context context.Context
}

func (p *UniversalEngineParams) GetContext() context.Context {
	if p.Context == nil {
		return context.Background()
	}
	return p.Context
}

// NewUniversalEngineParams constructs an initialized UniversalEngineParams.
// It automatically injects two event handlers into the Events system:
//   - CallEventsStartEvent - logs the start of any event call (debug level)
//   - CallEventsEndEvent   - logs the end, including any error
//
// Parameters: generator, events, scope, logger. All must be non‑nil.
// Returns a filled struct or an error if event registration fails.
func NewUniversalEngineParams(
	generator *Generator,
	events *Events,
	scope ScopeType,
	logger *Logger,
	context context.Context,
) (*UniversalEngineParams, error) {
	if generator == nil || events == nil || logger == nil {
		return nil, fmt.Errorf("Invalid input: nil refs")
	}
	logS := func(_name interface{}, _ *Events) error {
		name, ok := _name.(string)
		if !ok {
			return fmt.Errorf("Invalid event input: bad interface")
		}
		logger.PrintLog("debug", "Start call '"+name+"' event")
		return nil
	}
	logE := func(_name interface{}, _ *Events) error {
		name, ok := _name.(string)
		if !ok {
			return fmt.Errorf("Invalid event input: bad interface")
		}
		text := "End call '" + name + "' event"
		logger.PrintLog("debug", text)
		return nil
	}
	events.NewEvent(CallEventsStartEvent, logS)
	events.NewEvent(CallEventsEndEvent, logE)
	return &UniversalEngineParams{
		Generator: generator,
		Event:     events,
		Scope:     scope,
		Logger:    logger,
		Context:   context,
	}, nil
}
