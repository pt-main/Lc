package core

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type LoggerInterface interface {
	GetStatusForm(status string) string
	SetStatusForm(status, form string)
	PrintLog(status, message string)
	GetLog() string
}

// Logger is a thread-safe structured logger for engine diagnostics.
// It stores a list of log lines and supports a custom format per status.
// Typical usage: attach to UniversalEngineParams.Logger.
type Logger struct {
	mu                sync.RWMutex
	Logging           map[string]bool
	Log               []string
	MaxLogLength      int
	Statuses          map[string]string
	DefaultStatusForm string
}

// GetStatusForm returns the format string associated with the given status,
// falling back to DefaultStatusForm when the status has none.
func (l *Logger) GetStatusForm(status string) string {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if s, ok := l.Statuses[status]; ok {
		return s
	}
	return l.DefaultStatusForm
}

func (l *Logger) SetStatusForm(status, form string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Statuses[status] = form
}

// PrintLog appends a formatted line to the internal slice, and writes it to
// stdout when the status is enabled in Logging.
func (l *Logger) PrintLog(status string, message string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	format := l.DefaultStatusForm
	if s, ok := l.Statuses[status]; ok {
		format = s
	}
	line := fmt.Sprintf(format, status, time.Now().UTC(), message)
	if l.Logging[status] {
		fmt.Println(line)
	}
	l.Log = append(l.Log, line)
	if len(l.Log) > l.MaxLogLength && l.MaxLogLength > 0 {
		l.Log = l.Log[1:]
	}
}

// GetLog returns the retained log lines joined by newlines.
func (l *Logger) GetLog() string {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return strings.Join(l.Log, "\n")
}

// NewLogger creates a Logger with an optional defaultStatusForm. The format
// uses three placeholders: %s for status, %v for timestamp and %s for the
// message. An empty form selects the default "%s [%v] [%s]\n".
func NewLogger(defaultStatusForm string) *Logger {
	if defaultStatusForm == "" {
		defaultStatusForm = "%s [%v] [%s]\n"
	}
	return &Logger{
		Log:               make([]string, 0),
		Statuses:          make(map[string]string),
		DefaultStatusForm: defaultStatusForm,
		Logging:           make(map[string]bool),
		MaxLogLength:      -1,
	}
}
