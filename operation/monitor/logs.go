package monitor

import (
	api_operation "github.com/wunderkraut/radi-api/operation"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

/**
 * The logs monitor operations is about providing a serial
 * log output response attached to a writer.  There are no
 * readability requirements for this, as opposed to status
 * and info
 */

const (
	OPERATION_ID_MONITOR_LOGS = "monitor.logs"
)

// Base class for monitor info Operation
type BaseMonitorLogsOperation struct {
	MonitorBaseWriterOperation
}

// Id the operation
func (logs *BaseMonitorLogsOperation) Id() string {
	return OPERATION_ID_MONITOR_LOGS
}

// Label the operation
func (logs *BaseMonitorLogsOperation) Label() string {
	return "Log outputs"
}

// Description for the operation
func (logs *BaseMonitorLogsOperation) Description() string {
	return "Logging output information."
}

// Man page for the operation
func (logs *BaseMonitorLogsOperation) Help() string {
	return ""
}

// Is this an internal API operation
func (logs *BaseMonitorLogsOperation) Usage() api_usage.Usage {
	return api_operation.Usage_External()
}
