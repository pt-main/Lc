package core

const (
	StringParseEvent = "input string->parsed []ParsedNode"
	StringCallEvent  = "call(parsed []ParsedNode)"
	ByteParseEvent   = "input []byte->parsed []ParsedBytes"
	ByteCallEvent    = "call(parsed []ParsedBytes)"
)

type CommandType func([]interface{}) error

type CommandMeta struct {
	Handler CommandType
	Doc     string
}

type ScopeType map[string]interface{}

type EventType func(interface{}) error
