package core

import (
	"context"

	"github.com/pt-main/lc/v2/public"
	"github.com/pt-main/lc/v2/public/errors"
)

// UniversalEngineParams bundles what every command handler needs: the
// Generator for output, the Events manager for hooks, the shared Scope, the
// Logger and the active Context.
type UniversalEngineParams struct {
	Generator *Generator
	Event     EventsInterface
	Scope     ScopeType
	Logger    LoggerInterface
	Context   context.Context
}

// GetContext returns the active context, or context.Background() when none is
// set, so handlers need no nil check.
func (p *UniversalEngineParams) GetContext() context.Context {
	if p.Context == nil {
		return context.Background()
	}
	return p.Context
}

// NewUniversalEngineParams builds the params and registers the handlers of
// CallEventsStartEvent and CallEventsEndEvent, which trace event dispatch
// through the logger on the "event" level.
//
// Err errors.CorePackageSystemError: nil generator, events or logger.
func NewUniversalEngineParams(
	generator *Generator,
	events *Events,
	scope ScopeType,
	logger *Logger,
	ctx context.Context,
) (*UniversalEngineParams, ErrorInterface) {
	if generator == nil || events == nil || logger == nil {
		return nil, Err(errors.CorePackageSystemError, "Invalid input: nil refs")
	}
	logEvent := func(stage string) EventType {
		return func(e *Events, _ *EventInput) ErrorInterface {
			name, err := ScopeGetSynced[string](e.Scope(), public.EventsScopeCallName)
			if err != nil {
				return Wrap(errors.CorePackageSystemError, err, "LogEvent %s failed", stage)
			}
			logger.PrintLog("event", stage+" call '"+name+"' event")
			return nil
		}
	}
	events.NewEvent(public.CallEventsStartEvent, logEvent("Start"))
	events.NewEvent(public.CallEventsEndEvent, logEvent("End"))
	return &UniversalEngineParams{
		Generator: generator,
		Event:     events,
		Scope:     scope,
		Logger:    logger,
		Context:   ctx,
	}, nil
}
