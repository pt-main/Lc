package public

const (
	StringParseEvent = "INPUT string->PARSED []ParsedNode"
	StringCallEvent  = "call(PARSED []ParsedNode)"
	ByteParseEvent   = "INPUT []byte->PARSED []ParsedBytes"
	ByteCallEvent    = "call(PARSED []ParsedBytes)"
)

const (
	CallEventsStartEvent = "->call(Events.CallEvents)"
	CallEventsEndEvent   = "call(Events.CallEvents)->"
)
