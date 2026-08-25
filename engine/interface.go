package engine

import (
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/parsing/stringParsing"
)

type EngineInterface[CmdT int | string | byte | float32 | float64,
	ParserInput any, ParserOutput any] interface {
	Process(ParserInput) core.ErrorInterface
	NewCommand(CmdT, core.CommandType[EngineInterface[
		CmdT, ParserInput, ParserOutput], ParserOutput], *core.SimpleInput) error
	GetUep() *core.UniversalEngineParams
	GetParser() parsing.ParserInterface[ParserInput, ParserOutput]
	GetCommands() map[CmdT]core.CommandMeta[EngineInterface[
		CmdT, ParserInput, ParserOutput], ParserOutput]
}

type StringEngineInterface = EngineInterface[string, string, stringParsing.ParsedNode]
type ByteEngineInterface = EngineInterface[int, []byte, byteParsing.ParsedBytes]
