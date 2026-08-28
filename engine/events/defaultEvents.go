package events

import (
	"context"
)

type DefaultEvents struct{}

type CallLoopData[P, E any] struct {
	Ctx    context.Context
	Parsed []P
	Other  any
	Engine E
	Idx    *int
}
