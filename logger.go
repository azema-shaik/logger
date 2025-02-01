package main

import (
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

type logRecord struct {
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
func createRecord(message string, level int) logRecord {
	pc, file, lineNo, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	return logRecord{
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
	Name     string
	Level    int
	Handlers []Handler
}

func GetLogger(name string) *Logger {
	return &Logger{
		Name: name,
	}
}

func (l *Logger) SetLevel(level int) {
	l.Level = level
}

func (l *Logger) AddHandlers(handler Handler) {
	l.Handlers = append(l.Handlers, handler)
}

func (l *Logger) log(message string, level int) {
	if level < l.Level {
	}
	logRecord := createRecord(message, level)
	for _, hdlr := range l.Handlers {
		hdlr.emit(logRecord)
	}
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
