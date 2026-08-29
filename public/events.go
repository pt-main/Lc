package public

const (
	StringParseEvent       = "STRING:INPUT string->PARSED []ParsedNode"
	StringCallEvent        = "STRING:call(PARSED []ParsedNode)"
	StringCallCalloopEvent = "STRINGCALLOP call(PARSED []ParsedNode)"

	// Has no AstParseEvent, uses StringParseEvent
	AstCallCalloopEvent = "AST:CALLOP call(PARSED []ParsedNode)"
	AstCallEvent        = "AST:call(PARSED []ParsedNode)"
	AstCommandCallEvent = "AST:callCommand(PARSED []ParsedNode)" // with input = *AstCommandCtx

	ByteParseEvent       = "BYTE:INPUT []byte->PARSED []ParsedBytes"
	ByteCallEvent        = "BYTE:call(PARSED []ParsedBytes)"
	ByteCallHotloopEvent = "BYTE:HOTLOOP call(PARSED []ParsedBytes)"
)

const (
	CallEventsStartEvent = "->call(Events.CallEvents)"
	CallEventsEndEvent   = "call(Events.CallEvents)->"
)
