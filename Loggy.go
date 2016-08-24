package loggy

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/logutils"
)

// Create : Create a new logger
func Create(minLevel string) *Loggy {
	// Default filter
	reflog := log.New(os.Stdout, "", log.LstdFlags)
	loggyFilter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel(minLevel),
		Writer:   os.Stdout,
	}
	reflog.SetOutput(loggyFilter)

	logger := Loggy(*reflog)
	return &logger
	// TODO: Enable/disable filters via config or switch
}

// Loggy : entry point for logging functionality
type Loggy log.Logger

// Warn : write a warning to the log
func (l *Loggy) Warn(v ...interface{}) {
	lo := log.Logger(*l)
	(&lo).Printf("[WARN] %s", v...)
}

// Warnf : write a warning to the log
func (l *Loggy) Warnf(format string, v ...interface{}) {
	lo := log.Logger(*l)
	(&lo).Printf("[WARN] %s", fmt.Sprintf(format, v...))
}

// Error : write an error to the log
func (l *Loggy) Error(v ...interface{}) {
	lo := log.Logger(*l)
	(&lo).Printf("[ERROR] %s", v...)
}

// Errorf : write an error to the log
func (l *Loggy) Errorf(format string, v ...interface{}) {
	lo := log.Logger(*l)
	(&lo).Printf("[ERROR] %s", fmt.Sprintf(format, v...))
}

// Debug : write a debg message to the log
func (l *Loggy) Debug(v ...interface{}) {
	lo := log.Logger(*l)
	(&lo).Printf("[DEBUG] %s", v...)
}

// Debugf : write a debg message to the log
func (l *Loggy) Debugf(format string, v ...interface{}) {
	lo := log.Logger(*l)
	(&lo).Printf("[DEBUG] %s", fmt.Sprintf(format, v...))
}
