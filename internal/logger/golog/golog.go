package golog

import (
	"log"
	"os"

	li "github.com/SaYaku64/showcase/internal/logger"
)

type gologModule struct {
	info, error, warning, fatal *log.Logger
}

// Initialization of gologModule. Compares with ILogger interface
func Init() li.ILogger {
	logger := gologModule{
		info:    log.New(os.Stdout, "[INFO] ", log.Lshortfile),
		error:   log.New(os.Stderr, "[ERROR] ", log.Lshortfile),
		warning: log.New(os.Stderr, "[WARNING] ", log.Lshortfile),
		fatal:   log.New(os.Stderr, "[FATAL] ", log.Lshortfile),
	}

	return &logger
}

// Info - log for the common info
func (l *gologModule) Info(format string, v ...any) {
	format = newSpace(format)

	l.info.Printf(format, v...)
}

// Error - log for errors
func (l *gologModule) Error(format string, v ...any) {
	format = newSpace(format)

	l.error.Printf(format, v...)
}

// Warning - log for the critical errors
func (l *gologModule) Warning(format string, v ...any) {
	format = newSpace(format)

	l.warning.Printf(format, v...)
}

// Fatal - log for the fatal errors, that needs to restart service
func (l *gologModule) Fatal(format string, v ...any) {
	format = newSpace(format)

	l.fatal.Fatalf(format, v...)
}

// newSpace - checks if it is new line at the end of string
func newSpace(str string) string {
	sLen := len(str)

	if sLen < 2 {
		str += "\n"

		return str
	}

	if str[sLen-1:] != "\n" {
		str += "\n"
	}

	return str
}
