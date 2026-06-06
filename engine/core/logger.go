package core

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Logger struct {
	mu                sync.RWMutex
	Log               []string
	Statuses          map[string]string
	DefaultStatusForm string
}

func (l *Logger) GetStatusForm(status string) string {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if s, ok := l.Statuses[status]; ok {
		return s
	}
	return l.DefaultStatusForm
}

func (l *Logger) PrintLog(status string, message string) {
	format := l.GetStatusForm(status)
	line := fmt.Sprintf(format, status, time.Now().UTC(), message)
	fmt.Print(line)
	l.mu.Lock()
	l.Log = append(l.Log, line)
	l.mu.Unlock()
}

func (l *Logger) GetLog() string {
	return strings.Join(l.Log, "\n")
}

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
