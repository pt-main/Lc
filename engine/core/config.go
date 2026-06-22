package core

import "slices"

type CommandType[E, N any] func(*E, N) error

type CommandMeta[E, N any] struct {
	Handler CommandType[E, N]
	Doc     string
}

type EventType func(interface{}, *Events) error

type Option struct {
	Inited []int
	Flags  []string
	Scope  ScopeType
}

func (o *Option) HasFlag(name string) bool {
	return slices.Contains(o.Flags, name)
}
