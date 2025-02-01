package main

import "os"

// StreamHandler

type StreamHandler struct {
	writer    *os.File
	formatter Formatter
	logLevel  int
	filters   []Filter
}

func GetStreamHandler() StreamHandler {
	return StreamHandler{writer: os.Stdout}
}

func (s *StreamHandler) SetLevel(level int) {
	s.logLevel = level
}
func (s *StreamHandler) SetFormatter(formatter Formatter) {
	s.formatter = formatter
}
func (s *StreamHandler) AddFilter(filter Filter) {
	s.filters = append(s.filters, filter)
}

func (s *StreamHandler) filter(record logRecord) bool {

	if record.LevelNo < s.logLevel {
		return false
	}
	var isValidRecord int
	for _, filter := range s.filters {
		if filter.Filter(record) {
			isValidRecord += 1
		}
	}

	switch isValidRecord {
	case len(s.filters):
		return true
	default:
		return false
	}
}

func (s *StreamHandler) Emit(record logRecord) (int, error) {
	if !s.filter(record) {
		return 0, nil
	}
	message := s.formatter.Format(record)
	return s.writer.WriteString(message)
}

func (s *StreamHandler) Format(l logRecord) string {
	if s.formatter == nil {
		s.formatter = &StdFormatter{}
	}
	return s.formatter.Format(l)

}

//File Handler

// type FileHandler struct {
// 	writer *os.File
// 	formatter
// }
