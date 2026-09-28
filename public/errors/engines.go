package errors

const (
	ByteEngineProcessError1 ErrorCodeType = "Byte:PROCESS_ERR1"
	ByteEngineProcessError2 ErrorCodeType = "Byte:PROCESS_ERR2"

	StringEngineProcessError1 ErrorCodeType = "String:PROCESS_ERR1"
	StringEngineProcessError2 ErrorCodeType = "String:PROCESS_ERR2"

	AstEngineProcessError1 ErrorCodeType = "Ast:PROCESS_ERR1"
	AstEngineProcessError2 ErrorCodeType = "Ast:PROCESS_ERR2"
	AstEngineHandlerError  ErrorCodeType = "Ast:HANDLER_ERROR"
	AstEngineUnknown       ErrorCodeType = "Ast:UNKNOWN_NODE"

	EngineLifecycleEnd ErrorCodeType = "Engine:LIFECYCLE_END"
)

const (
	DefaultEventsSystemError          ErrorCodeType = "SYSTEM@DEFAULT_EVENTS"
	DefaultEventsPanicError           ErrorCodeType = "DEFAULT_EVENTS:PANIC"
	DefaultEventsCallErrorCmdNotFound ErrorCodeType = "DEFAULT_EVENTS:CMD_NOT_FOUND"
	DefaultEventsCallErrorContexted   ErrorCodeType = "DEFAULT_EVENTS:CONTEXT_ERROR"
	DefaultEventsCallErrorHandler     ErrorCodeType = "DEFAULT_EVENTS:HANDLER"
	DefaultEventsCallErrorUnknown     ErrorCodeType = "DEFAULT_EVENTS:UNKNOWN"
)
