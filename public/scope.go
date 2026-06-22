package public

const (
	ByteEngineScopeParsed      = "PARSED []ParsedBytes"
	ByteEngineScopeEndianess   = "ENDIANESS int"
	ByteEngineScopeBytecodeIdx = "BYTECODE_IDX *int"
	ByteEngineScopeInput       = "INPUT []byte"
)

const (
	StringEngineScopeInput  = "INPUT string"
	StringEngineScopeParsed = "PARSED []ParsedNode"
)

const (
	EventsScopeCallName  = "CALL_NAME string"
	EventsScopeCallError = "CALL_ERROR error"
)

const (
	PluginsScopeEuPtr = "EuPtr *EngineUniversal"
	EuScopePmPtr      = "PmPtr *PluginManager"
)
