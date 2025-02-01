package main

import (
	"os"
)

func filter(handler Handler, record logRecord) bool {

	if record.LevelNo < (handler).GetLogLevel() {
		return false
	} else if len(handler.GetFilters()) == 0 {
		return true
	}

	var isValidRecord int
	for _, filter := range (handler).GetFilters() {
		if filter.Filter(record) {
			isValidRecord += 1
		}
	}

	switch isValidRecord {
	case len((handler).GetFilters()):
		return true
	default:
		return false
	}
}

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

func (s *StreamHandler) SetLogLevel(level int) {
	s.logLevel = level
}

func (s *StreamHandler) GetLogLevel() int {
	return s.logLevel
}

func (s *StreamHandler) SetFormatter(formatter Formatter) {
	s.formatter = formatter
}
func (s *StreamHandler) AddFilter(filter Filter) {
	s.filters = append(s.filters, filter)
}

func (s *StreamHandler) GetFilters() []Filter {
	return (s).filters
}

func (s *StreamHandler) emit(record logRecord) (int, error) {
	if !filter(s, record) {
		return 0, nil

	}
	message := s.formatter.Format(record)
	return s.writer.WriteString(message)
}

func (s *StreamHandler) format(l logRecord) string {
	if s.formatter == nil {
		s.formatter = &StdFormatter{}
	}
	return s.formatter.Format(l)

}

func (s *StreamHandler) Close() {
	s.writer.Close()
}

//File Handler

type FileHandler struct {
	writer    *os.File
	formatter Formatter
	logLevel  int
	filters   []Filter
}

func GetFileHandler(filename string, flag int, perm os.FileMode) FileHandler {
	file, _ := os.OpenFile(filename, flag, perm)
	return FileHandler{writer: file}
}

func (f *FileHandler) SetLogLevel(level int) {
	f.logLevel = level
}

func (f *FileHandler) GetLogLevel() int {
	return f.logLevel
}

func (f *FileHandler) SetFormatter(formatter Formatter) {
	f.formatter = formatter
}

func (f *FileHandler) AddFilter(filter Filter) {
	f.filters = append(f.filters, filter)
}

func (f *FileHandler) GetFilters() []Filter {
	return f.filters
}

func (f *FileHandler) emit(record logRecord) (int, error) {
	if !filter(f, record) {
		return 0, nil

	}
	message := f.formatter.Format(record)
	return f.writer.WriteString(message)
}

func (f *FileHandler) format(l logRecord) string {
	if f.formatter == nil {
		f.formatter = &StdFormatter{}
	}
	return f.formatter.Format(l)
}

func (f *FileHandler) Close() {
	f.writer.Close()
}
