package dd

import (
	"log"
	"os"
)

// A Logger is a minimalistic interface to log messages to. Should
// be used to provide custom logging writers  use.
type Logger interface {
	Log(...any)
}

// A LoggerFunc is a convenience type to convert a function taking a variadic
// list of arguments and wrap it so the Logger interface can be used.
type LoggerFunc func(...any)

// Log calls the wrapped function with the arguments provided
func (f LoggerFunc) Log(args ...any) {
	f(args...)
}

type LogLevel int

const (
	ERROR LogLevel = iota
	WARN
	INFO
	DEBUG
)

// A LevelLogger is a interface to log message with level and can be filter by logger
type LevelLogger interface {
	Logger
	Error(...any)
	Warn(...any)
	Info(...any)
	Debug(...any)
}

// NewDefaultLogger returns a Logger which will write log messages to stdout, and
// use same formatting runes as the stdlib log.Logger
func NewDefaultLogger() Logger {
	return &defaultLogger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// A defaultLogger provides a minimalistic logger satisfying the Logger interface.
type defaultLogger struct {
	logger *log.Logger
}

// Log logs the parameters to the stdlib logger. See log.Println.
func (l defaultLogger) Log(args ...any) {
	l.logger.Println(args...)
}

// NewLevelLogger returns a LevelLogger which will write log to stdout, and filter
// with LogLevel
func NewLevelLogger(level LogLevel) LevelLogger {
	return &defaultLevelLogger{
		errorLogger: log.New(os.Stderr, "E", log.LstdFlags|log.Lmicroseconds|log.Lshortfile|log.Lmsgprefix),
		warnLogger:  log.New(os.Stdout, "W", log.LstdFlags|log.Lmicroseconds|log.Lshortfile|log.Lmsgprefix),
		infoLogger:  log.New(os.Stdout, "I", log.LstdFlags|log.Lmicroseconds|log.Lshortfile|log.Lmsgprefix),
		debugLogger: log.New(os.Stdout, "D", log.LstdFlags|log.Lmicroseconds|log.Llongfile|log.Lmsgprefix),
		level:       level,
	}
}

type defaultLevelLogger struct {
	errorLogger *log.Logger
	warnLogger  *log.Logger
	infoLogger  *log.Logger
	debugLogger *log.Logger
	level       LogLevel
}

func (l defaultLevelLogger) Log(args ...any) {
	l.Info(args...)
}

func (l defaultLevelLogger) Error(args ...any) {
	if l.level >= ERROR {
		l.errorLogger.Println(args...)
	}
}

func (l defaultLevelLogger) Warn(args ...any) {
	if l.level >= WARN {
		l.warnLogger.Println(args...)
	}
}

func (l defaultLevelLogger) Info(args ...any) {
	if l.level >= INFO {
		l.infoLogger.Println(args...)
	}
}

func (l defaultLevelLogger) Debug(args ...any) {
	if l.level >= DEBUG {
		l.debugLogger.Println(args...)
	}
}
