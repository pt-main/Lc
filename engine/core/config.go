package core

const (
	StringParseEvent = "input string->parsed []ParsedNode"
	StringCallEvent  = "call(parsed []ParsedNode)"
	ByteParseEvent   = "input []byte->parsed []ParsedBytes"
	ByteCallEvent    = "call(parsed []ParsedBytes)"
)

type CommandType[E, N any] func(*E, N) error

type CommandMeta[E, N any] struct {
	Handler CommandType[E, N]
	Doc     string
}

type ScopeType map[string]interface{}

type EventType func(interface{}, *Events) error
