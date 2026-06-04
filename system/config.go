package system

const (
	StringParseEvent = "input_string->parsed_[]ParsedNode"
	StringCallEvent  = "call(parsed_[]ParsedNode)"
	ByteParseEvent   = "input_[]byte->parsed_[]ParsedBytes"
	ByteCallEvent    = "call(parsed_[]ParsedBytes)"
)

type CommandType func([]interface{}) error

type CommandMeta struct {
	Handler CommandType
	Doc     string
}

type ScopeType map[string]interface{}

type EventType func(interface{}) error
