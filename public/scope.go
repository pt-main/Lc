package public

const (
	ByteEngineScopeParsed      = "PARSED []ParsedBytes"
	ByteEngineScopeEndianess   = "ENDIANESS int"
	ByteEngineScopeBytecodeIdx = "BYTECODE_IDX *int"
	ByteEngineScopeInput       = "INPUT []byte"
)

const (
	StringEngineScopeInput    = "INPUT string"
	StringEngineScopeParsed   = "PARSED []ParsedNode"
	StringEngineScopeInstrIdx = "INSTR_IDX *int"
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
