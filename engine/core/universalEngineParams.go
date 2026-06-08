package core

import "fmt"

// UniversalEngineParams is a container shared by both StringEngine and ByteEngine.
// It holds the Generator (for output accumulation), Events (for hooks),
// Scope (for passing arbitrary data between stages), and Logger (for diagnostics).
// This struct is embedded (not composed by pointer) in the engine types,
// promoting its fields and methods to the engine itself.
type UniversalEngineParams struct {
	// Generator *Generator – controls code/output generation across pipeline points.
	Generator *Generator
	// Event *Events – allows hooking into parsing and command dispatch.
	Event *Events
	// Scope ScopeType – a map[string]interface{} that can be used to share
	//   variables between events, parsers, and command handlers.
	Scope ScopeType
	// Logger *Logger – if set, logs internal steps (event calls, errors).
	Logger *Logger
}

// NewUniversalEngineParams constructs an initialized UniversalEngineParams.
// It automatically injects two event handlers into the Events system:
//   - CallEventsStartEvent – logs the start of any event call (debug level)
//   - CallEventsEndEvent   – logs the end, including any error
//
// Parameters: generator, events, scope, logger. All must be non‑nil.
// Returns a filled struct or an error if event registration fails.
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
