package system

const (
	ParseEvent = "input_string->parsed_[]ParsedNode"
	CallEvent  = "call(parsed_[]ParsedNode)"
)

type CommandType func([]interface{}) error

type CommandMeta struct {
	Handler CommandType
	Doc     string
}

type ScopeType map[string]interface{}

type EventType func(interface{}) error
