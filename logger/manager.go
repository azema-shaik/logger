package logger

import (
	"strings"
)

type placeholder struct {
	loggerMap []*Logger
}

func (p *placeholder) append(logger *Logger) {
	p.loggerMap = append(p.loggerMap, logger)
}

type Manager struct {
	rootLogger *Logger
	loggerDict map[string]LoggerLike
}

func (m *Manager) GetLogger(name string) *Logger {
	//check if the logger with same name exisits
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
			logger = &Logger{Name: name, Propagate: true}
			logger.manager = m
			m.fixUpChildren(ph, logger) // reset its children parent attribyte to itself
			m.fixUpParents(logger)

		}
	} else { //this name does not exist
		logger = &Logger{Name: name, Propagate: true}
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
			m.loggerDict[substr] = &placeholder{}
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
		if _logger.parent.Name[:len(name)] != name {
			_logger.parent.Name = name
		}
	}

}

// func (m *Manager) fixUpChildren(logger *Logger) {
// 	name := logger.Name
// 	findDot := strings.LastIndex(name, ".")

// 	var rv *Logger
// 	for (findDot > 0) && (rv == nil) {
// 		substr := name[:findDot]
// 		// this condition checks is the parent with given exists
// 		// if it doesnot exist an empty logger is initialized
// 		if m.loggerDict[name] == nil {
// 			var emptyLogger *Logger
// 			m.loggerDict[substr] = emptyLogger
// 		} else {
// 			parent := m.loggerDict[substr]
// 			if parent != nil {
// 				rv = parent
// 			} else {

// 			}
// 		}

// 	}
// }
