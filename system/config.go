package system

import "github.com/pt-main/lc/parsing"

const (
	ParseEvent = "input_string->parsed_[]ParsedNode"
	CallEvent  = "call(parsed_[]ParsedNode)"
)

type CommandType func(*Engine, parsing.ParsedNode) error

type CommandMeta struct {
	Handler CommandType
	Doc     string
}

type ScopeType map[string]interface{}

type EventType func(*Engine) error
