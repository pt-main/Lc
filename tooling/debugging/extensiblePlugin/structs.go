package extensiblePlugin

import (
	"context"

	"github.com/pt-main/lc/v2/engine"
	"github.com/pt-main/lc/v2/engine/events"
	"github.com/pt-main/lc/v2/parsing/stringParsing"
)

type CLEData[I, P, E any] struct {
	Input  I
	Idx    *int
	Parsed []P
	PLen   int
	Ctx    context.Context
	E      E
}

type SCLEData CLEData[
	events.StringCLDType,
	stringParsing.ParsedNode,
	engine.StringEngineInterface,
]

type BCLEData CLEData[
	events.ByteCLDType,
	events.ByteCallAttr,
	engine.ByteEngineInterface,
]
