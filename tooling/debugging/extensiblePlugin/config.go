package extensiblePlugin

// Call loop events
const (
	CLEPreEvent    = "CallLoopE PreEvent"
	CLEInPreEvent  = "CallLoopE InPreEvent"
	CLEInPostEvent = "CallLoopE InPostEvent"
	CLEPostEvent   = "CallLoopE PostEvent"
)

const (
	CLEScopeData = "ExtensiblePlugin ScopeData CallLoopE Data" // Call loop data (CLEData)
)

const (
	ECLFlag = "ExtensiblePlugin" // This flag will be added to plugins manager when the plugin is initialized
)
