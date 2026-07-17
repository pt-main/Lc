package extensiblePlugin

import (
	"context"

	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/events"
	"github.com/pt-main/lc/parsing/stringParsing"
)

type CLEData[I, P, E any] struct {
	Input  I
	Idx    *int
	Parsed []P
	PLen   int
	Ctx    context.Context
	E      *E
}

type SCLEData CLEData[
	events.StringCLDType,
	stringParsing.ParsedNode,
	engine.StringEngine,
]

type BCLEData CLEData[
	events.ByteCLDType,
	events.ByteCallAttr,
	engine.ByteEngine,
]
