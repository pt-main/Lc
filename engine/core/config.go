package core

import (
	"slices"
)

type CommandType[E, N any] func(*E, N) error

type CommandMeta[E, N any] struct {
	Handler CommandType[E, N]
	Doc     string
}

type EventType func(*Events, *EventInput) error

type Option struct {
	Flags []string
	Scope ScopeType
}

func (o *Option) HasFlag(name string) bool {
	return slices.Contains(o.Flags, name)
}

type EventInput struct {
	Option *Option
	Input  interface{}
}
