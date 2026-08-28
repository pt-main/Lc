package public

const (
	ByteEngineScopeParsed                = "PARSED []ParsedBytes"
	ByteEngineScopeEndianess             = "ENDIANESS int"
	ByteEngineScopeBytecodeIdx           = "BYTECODE_IDX *int"
	ByteEngineScopeHotloopCtxCheckPeriod = "CTX_CKECK_PERIOD int"
	ByteEngineScopeInput                 = "INPUT []byte"
)

const (
	StringEngineScopeInput        = "INPUT string"
	StringEngineScopeParsed       = "PARSED []ParsedNode"
	StringEngineScopeInstrIdx     = "INSTR_IDX *int"
	StringEngineScopeCanBeUnknown = "StringEngineScopeCanBeUnknown bool"
)

const (
	EventsScopeCallName  = "CALL_NAME string"
	EventsScopeCallError = "CALL_ERROR error"
	EventsScopeDERawLine = "RAW_LINE string"
)

const (
	PluginsScopeEuPtr = "EuPtr *EngineUniversal"
	EuScopePmPtr      = "PmPtr *PluginManager"
)
