package errors

const (
	ByteEngineProcessError1   ErrorCodeType = "Byte:PROCESS_ERR1"
	ByteEngineProcessError2   ErrorCodeType = "Byte:PROCESS_ERR2"
	StringEngineProcessError1 ErrorCodeType = "String:PROCESS_ERR1"
	StringEngineProcessError2 ErrorCodeType = "String:PROCESS_ERR2"
)

const (
	DefaultEventsSystemError          ErrorCodeType = "SYSTEM@DEDAULT_EVENTS"
	DefaultEventsPanicError           ErrorCodeType = "DEDAULT_EVENTS:PANIC"
	DefaultEventsCallErrorCmdNotFound ErrorCodeType = "DEDAULT_EVENTS:CMD_NOT_FOUND"
	DefaultEventsCallErrorContexted   ErrorCodeType = "DEDAULT_EVENTS:CONTEXTED_ERROR"
	DefaultEventsCallErrorContex      ErrorCodeType = "DEDAULT_EVENTS:CONTEXT_ERROR"
	DefaultEventsCallErrorHandler     ErrorCodeType = "DEDAULT_EVENTS:HANDLER"
	DefaultEventsCallErrorUnknown     ErrorCodeType = "DEDAULT_EVENTS:UNKNOWN"
)
