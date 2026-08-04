package extensiblePlugin

// Сalloop events
const (
	CLEPreEvent    = "CalloopE PreEvent"
	CLEInPreEvent  = "CalloopE InPreEvent"
	CLEInPostEvent = "CalloopE InPostEvent"
	CLEPostEvent   = "CalloopE PostEvent"
)

const (
	CLEScopeData = "ExtensiblePlugin ScopeData CalloopE Data" // Сalloop data (CLEData)
)

const (
	ECLFlag = "ExtensiblePlugin" // This flag will be added to plugins manager when the plugin is initialized
)
