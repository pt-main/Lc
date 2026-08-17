package core

import (
	"slices"
)

type CommandType[EI, N any] func(EI, *N) ErrorInterface

type CommandMeta[EI, N any] struct {
	Handler CommandType[EI, N]
	Doc     string
}

type EventType func(*Events, *EventInput) ErrorInterface

type Option struct {
	Flags []string
	Scope ScopeType
}

func (o *Option) HasFlag(name string) bool {
	return slices.Contains(o.Flags, name)
}

type SimpleInput struct {
	Option *Option
	Input  interface{}
}

type EventInput = SimpleInput
