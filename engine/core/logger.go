package core

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Logger is a thread-safe structured logger for engine diagnostics.
// It stores a list of log lines, supports custom status formats, and
// allows different formatting per status (e.g., "error", "info", "debug").
// Typical usage: attach to UniversalEngineParams.Logger.
type Logger struct {
	mu                sync.RWMutex
	Log               []string
	Statuses          map[string]string
	DefaultStatusForm string
}

// GetStatusForm returns the format string associated with the given status,
// falling back to DefaultStatusForm if no custom format is set.
// Custom formats can be registered by directly assigning to l.Statuses map.
func (l *Logger) GetStatusForm(status string) string {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if s, ok := l.Statuses[status]; ok {
		return s
	}
	return l.DefaultStatusForm
}

// PrintLog writes a log entry to stdout and appends it to the internal slice.
// The status string determines the format via GetStatusForm (if a custom
// format for that status exists). The message is inserted into the format.
// Example: logger.PrintLog("error", "failed to parse token")
func (l *Logger) PrintLog(status string, message string) {
	format := l.GetStatusForm(status)
	line := fmt.Sprintf(format, status, time.Now().UTC(), message)
	fmt.Print(line)
	l.mu.Lock()
	l.Log = append(l.Log, line)
	l.mu.Unlock()
}

// GetLog returns the entire log as a single string with newline separators.
// It is useful for saving logs to a file or showing them after execution.
func (l *Logger) GetLog() string {
	return strings.Join(l.Log, "\n")
}

// NewLogger creates a new Logger with an optional defaultStatusForm.
// The format uses three placeholders: %v for status, %v for timestamp,
// and %s for the message. Example default: "%v [%v] [%s]"
// If empty string is passed, the default format is used.
func NewLogger(defaultStatusForm string) *Logger {
	if defaultStatusForm == "" {
		defaultStatusForm = "%v [%v] [%s]"
	}
	return &Logger{
		Log:               make([]string, 0),
		Statuses:          make(map[string]string),
		DefaultStatusForm: defaultStatusForm,
	}
}
