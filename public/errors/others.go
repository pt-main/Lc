package errors

const (
	ScopeGetError ErrorCodeType = "SCOPE_GET" /* metadata ScopeGetErrorMetakey : string */
)

const (
	CorePackageSystemError ErrorCodeType = "SYSTEM@CORE"
	CorePackageLcError     ErrorCodeType = "SYSTEM@LC"
	WrappedError           ErrorCodeType = "WrappedError"

	CorePackageLcLifecycleError ErrorCodeType = "SYSTEM@LC:LIFECYCLE"
)

const (
	ParsingError ErrorCodeType = "PARSING_ERROR"
)

const (
	BytecodeShiftError ErrorCodeType = "BYTECODE_SHIFT_ERROR"
)

const (
	ExtensiblePluginError ErrorCodeType = "EXTENSIBLE_PLUGIN_ERROR"
)

var ErrExit ErrorCodeType = "EXIT"
