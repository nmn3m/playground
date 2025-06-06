package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

var (
	infoColor    = color.New(color.FgGreen)
	warnColor    = color.New(color.FgYellow)
	errorColor   = color.New(color.FgRed)
	debugColor   = color.New(color.FgCyan)
	successColor = color.New(color.FgGreen, color.Bold)
	enableDebug  = false // Flag to enable/disable debug messages

)

func init() {
	enableDebug = os.Getenv("LOG_LEVEL") == "debug"
}

// Info prints info message with format
func Info(format string, args ...interface{}) {
	_, _ = infoColor.Printf(format+"\n", args...)
}

// Infof is an alias for Info for consistency
func Infof(format string, args ...interface{}) {
	Info(format, args...)
}

// Infoln prints info message with newline
func Infoln(format string, args ...interface{}) {
	_, _ = infoColor.Printf(format+"\n", args...)
}

// Warn prints warning message with format
func Warn(format string, args ...interface{}) {
	_, _ = warnColor.Printf(format+"\n", args...)
}

// Warnf is an alias for Warn for consistency
func Warnf(format string, args ...interface{}) {
	Warn(format, args...)
}

// Warnln prints warning message with newline
func Warnln(format string, args ...interface{}) {
	_, _ = warnColor.Printf(format+"\n", args...)
}

// Error prints error message with format (suppressed in silent mode)
func Error(format string, args ...interface{}) {
	_, _ = errorColor.Printf(format+"\n", args...)
}

// Errorf is an alias for Error for consistency
func Errorf(format string, args ...interface{}) {
	Error(format, args...)
}

// Errorln prints error message with newline (suppressed in silent mode)
func Errorln(format string, args ...interface{}) {
	_, _ = errorColor.Printf(format+"\n", args...)
}

// Debug prints debug message with format (suppressed in silent mode)
func Debug(format string, args ...interface{}) {
	if enableDebug {
		_, _ = debugColor.Printf(format+"\n", args...)
	}
}

// Debugf is an alias for Debug for consistency
func Debugf(format string, args ...interface{}) {
	if enableDebug {
		Debug(format, args...)
	}
}

// Debugln prints debug message with newline (suppressed in silent mode)
func Debugln(format string, args ...interface{}) {
	if enableDebug {
		_, _ = debugColor.Printf(format+"\n", args...)
	}
}

// Success prints success message with format
func Success(format string, args ...interface{}) {
	_, _ = successColor.Printf(format+"\n", args...)
}

// Successf is an alias for Success for consistency
func Successf(format string, args ...interface{}) {
	Success(format, args...)
}

// Successln prints success message with newline
func Successln(format string, args ...interface{}) {
	_, _ = successColor.Printf(format+"\n", args...)
}

// Fatal prints error message and exits
func Fatal(format string, args ...interface{}) {
	_, _ = errorColor.Printf(format+"\n", args...)
	os.Exit(1)
}

// Print prints plain message with format
func Print(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

// Println prints plain message with newline
func Println(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

// GetWriter returns an io.Writer for use with external libraries
func GetWriter() io.Writer {
	return os.Stdout
}
