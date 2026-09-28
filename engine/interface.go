package engine

import (
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/parsing/stringParsing"
)

// EngineInterface is the contract every engine backend satisfies. CmdT is the
// command key type, so the same interface covers string names and byte opcodes.
type EngineInterface[CmdT int | string | byte | float32 | float64,
	ParserInput any, ParserOutput any] interface {
	Process(ParserInput) core.ErrorInterface
	NewCommand(CmdT, core.CommandType[EngineInterface[
		CmdT, ParserInput, ParserOutput], ParserOutput], *core.SimpleInput) error
	GetUep() *core.UniversalEngineParams
	GetParser() parsing.ParserInterface[ParserInput, ParserOutput]
	GetCommands() map[CmdT]core.CommandMeta[EngineInterface[
		CmdT, ParserInput, ParserOutput], ParserOutput]
	// GetCommand looks up a single command without copying the whole map.
	GetCommand(CmdT) (core.CommandMeta[EngineInterface[
		CmdT, ParserInput, ParserOutput], ParserOutput], bool)
}

type StringEngineInterface = EngineInterface[string, string, stringParsing.ParsedNode]
type ByteEngineInterface = EngineInterface[int, []byte, byteParsing.ParsedBytes]
type AstEngineInterface = StringEngineInterface
