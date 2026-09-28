package errors

const (
	// ScopeGetError carries the offending key in EMK(0, "string").
	ScopeGetError ErrorCodeType = "SCOPE_GET"
)

const (
	CorePackageSystemError      ErrorCodeType = "SYSTEM@CORE"
	CorePackageLcError          ErrorCodeType = "SYSTEM@LC"
	CorePackageLcLifecycleError ErrorCodeType = "SYSTEM@LC:LIFECYCLE"
	WrappedError                ErrorCodeType = "WrappedError"
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
