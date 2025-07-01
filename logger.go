package logger

import (
	"fmt"
	"runtime"
	"sync"
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
	mu        sync.RWMutex
}

func (l *Logger) String() string {
	return fmt.Sprintf("<Logger(%#v, %s)>", l.Name, levelName(l.Level))
}

func GetLogger(name string) *Logger {
	if root == nil {
		root = &Logger{
			Name: "root", Level: WARNING,
		}
		root.manager = &Manager{rootLogger: root}
	}

	if len(name) == 0 || name == "root" {
		return root
	} else {
		return root.manager.GetLogger(name)
	}
}

func (l *Logger) GetParent() *Logger {
	return l.parent
}

func (l *Logger) GetManager() *Manager {
	return l.manager
}

func (l *Logger) SetLevel(level int) {
	l.Level = level
}

func (l *Logger) AddHandler(handler Handler) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.handlers = append(l.handlers, handler)
}

func (l *Logger) log(message string, level int) {
	if level < l.Level {
		return
	}
	logRecord := createRecord(l.Name, message, level)
	l.callHandlers(logRecord)
}

func (l *Logger) callHandlers(record LogRecord) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	logger := l
	found := 0
	for logger != nil {
		for _, hdlr := range logger.handlers {
			found += 1
			if _, err := hdlr.emit(record); err != nil {
				fmt.Printf("Error emitting log record: %v\n", err)
			}
		}

		if !logger.Propagate {
			logger = nil
		} else {
			logger = logger.parent

		}
	}

	if found == 0 {
		fmt.Printf("%s has no handlers", l.Name)

	}

}

func (l *Logger) Debug(message string) {
	l.log(message, DEBUG)
}

func (l *Logger) Info(message string) {
	l.log(message, INFO)
}

func (l *Logger) Error(message string) {
	l.log(message, ERROR)
}

func (l *Logger) Warning(message string) {
	l.log(message, WARNING)
}
func (l *Logger) Critical(message string) {
	l.log(message, CRITICAL)
}

func (l *Logger) Close() {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, logger := range root.manager.loggerDict {
		value, ok := logger.(*Logger)
		if !ok {
			continue
		}
		for _, hdlr := range value.handlers {
			hdlr.Close()
		}
	}
}
