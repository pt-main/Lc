package core

const (
	StringParseEvent = "INPUT string->PARSED []ParsedNode"
	StringCallEvent  = "call(PARSED []ParsedNode)"
	ByteParseEvent   = "INPUT []byte->PARSED []ParsedBytes"
	ByteCallEvent    = "call(PARSED []ParsedBytes)"
)

type CommandType[E, N any] func(*E, N) error

type CommandMeta[E, N any] struct {
	Handler CommandType[E, N]
	Doc     string
}

type ScopeType map[string]interface{}

type EventType func(interface{}, *Events) error
