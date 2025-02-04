package logger

import (
	"fmt"
	"runtime"
	"time"
)

const (
	_ = iota * 10
	DEBUG
	INFO
	ERROR
	WARNING
	CRITICAL
)

type LogRecord struct {
	Name      string
	Datetime  time.Time
	File      string
	LineNo    int
	Message   string
	FuncName  string
	LevelName string
	LevelNo   int
}

func levelName(level int) (levelname string) {
	switch level {
	case DEBUG:
		levelname = "DEBUG"
	case INFO:
		levelname = "INFO"
	case ERROR:
		levelname = "ERROR"
	case WARNING:
		levelname = "WARNING"
	case CRITICAL:
		levelname = "CRITICAL"
	}
	return levelname
}
func createRecord(name string, message string, level int) LogRecord {
	pc, file, lineNo, _ := runtime.Caller(3)
	funcName := runtime.FuncForPC(pc).Name()
	return LogRecord{
		Name:      name,
		Datetime:  time.Now(),
		File:      file,
		LineNo:    lineNo,
		Message:   message,
		LevelNo:   level,
		LevelName: levelName(level),
		FuncName:  funcName,
	}
}

type Logger struct {
	Propagate bool
	Name      string
	manager   *Manager
	Level     int
	handlers  []Handler
	parent    *Logger
}

func GetLogger(name string) *Logger {
	root = &Logger{
		Name: "root",
	}

	root.manager = &Manager{rootLogger: root}

	if name == "" && name == "root" {
		return root
	} else {
		return root.manager.GetLogger(name)
	}
}

func (l *Logger) SetLevel(level int) {
	l.Level = level
}

func (l *Logger) AddHandlers(handler Handler) {
	l.handlers = append(l.handlers, handler)
}

func (l *Logger) log(message string, level int) {
	if level < l.Level {
		return
	}
	logRecord := createRecord(l.Name, message, level)
	l.callHandlers(logRecord)
}

func (l *Logger) callHandlers(record LogRecord) (nBytes int, err error) {
	logger := l
	found := 0

	for logger != nil {
		for _, hdlr := range logger.handlers {
			found += 1
			if nBytes, err := hdlr.emit(record); err != nil {
				return nBytes, err
			}

			if !logger.Propagate {
				logger = nil
			} else {
				logger = logger.parent
			}
		}
	}

	if found == 0 {
		return nBytes, fmt.Errorf("%s has no handlers", l.Name)

	}

	return nBytes, err

}

func (l *Logger) Debug(message string) {
	l.log(message, DEBUG)
}

func (l *Logger) Info(message string) {
	l.log(message, INFO)
}

func (l *Logger) Warning(message string) {
	l.log(message, WARNING)
}
func (l *Logger) Critical(message string) {
	l.log(message, CRITICAL)
}

func (l *Logger) Close() {
	for _, hdlr := range l.handlers {
		hdlr.Close()
	}
}
