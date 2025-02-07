package logger

import (
	"strings"
	"sync"
)

type placeholder struct {
	loggerMap []*Logger
}

func (p *placeholder) GetLoggerMap() []*Logger {
	return p.loggerMap
}

func getPlaceHolder(logger *Logger) *placeholder {
	return &placeholder{
		loggerMap: []*Logger{logger},
	}

}

func (p *placeholder) append(logger *Logger) {
	var found bool
	for _, alogger := range p.loggerMap {
		if alogger == logger {
			found = true
		}
	}
	if !found {
		p.loggerMap = append(p.loggerMap, logger)
	}

}

type Manager struct {
	rootLogger *Logger
	loggerDict map[string]LoggerLike
	mu         *sync.Mutex
}

func (m *Manager) GetLoggerDict() map[string]LoggerLike {
	return m.loggerDict
}

func (m *Manager) GetLogger(name string) *Logger {
	//check if the logger with same name exisits
	m.mu.Lock()
	defer m.mu.Unlock()
	var logger *Logger
	if m.loggerDict == nil {
		m.loggerDict = make(map[string]LoggerLike)
	}

	if m.loggerDict[name] != nil {
		// if the logger already exists return the logger
		switch t := m.loggerDict[name].(type) {
		case *Logger:
			logger = t
		case *placeholder: //check if an empty placholder exists
			ph := t
			logger = &Logger{Name: name, Propagate: true, Level: WARNING}
			logger.manager = m
			m.fixUpChildren(ph, logger) // reset its children parent attribyte to itself
			m.fixUpParents(logger)

		}
	} else { //this name does not exist
		logger = &Logger{Name: name, Propagate: true, Level: WARNING}
		logger.manager = m
		m.loggerDict[name] = logger
		m.fixUpParents(logger)

	}
	return logger
}

func (m *Manager) fixUpParents(logger *Logger) {
	name := logger.Name
	findLastDot := strings.LastIndex(name, ".")
	var parent *Logger
	for findLastDot > 0 && parent == nil {
		substr := name[:findLastDot]
		if m.loggerDict[substr] == nil {
			m.loggerDict[substr] = getPlaceHolder(logger)
		} else {
			_parent := m.loggerDict[substr]
			switch t := _parent.(type) {
			case *Logger:
				parent = t
			case *placeholder:
				t.append(logger)
			}
		}
		findLastDot = strings.LastIndex(name[:findLastDot], ".")

	}

	if parent == nil {
		parent = m.rootLogger

	}

	logger.parent = parent

}

func (m *Manager) fixUpChildren(ph *placeholder, logger *Logger) {
	name := logger.Name
	for _, _logger := range ph.loggerMap {
		if !strings.HasPrefix(_logger.parent.Name, name) {
			_logger.parent.Name = name
		}
	}

}
