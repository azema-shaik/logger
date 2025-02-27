package logger

import (
	"os"
	"sync"
)

type BaseHandler struct {
	writer    *os.File
	formatter Formatter
	logLevel  int
	filters   []Filter
	mu        sync.RWMutex
}

func (b *BaseHandler) SetLogLevel(level int) {
	b.logLevel = level
}

func (b *BaseHandler) GetLogLevel() int {
	return b.logLevel
}

func (b *BaseHandler) SetFormatter(formatter Formatter) {
	b.formatter = formatter
}

func (b *BaseHandler) AddFilter(filter Filter) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.filters = append(b.filters, filter)
}

func (b *BaseHandler) GetFilters() []Filter {
	return b.filters
}

func (b *BaseHandler) emit(record LogRecord) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if record.LevelNo < b.logLevel || (len(b.filters) != 0 && !b.filter(record)) {
		return 0, nil

	}

	if b.formatter == nil {
		b.formatter = &StdFormatter{}
	}

	message := b.formatter.Format(record)
	nBytes, err := b.writer.WriteString(message)
	if err == nil {
		if b.writer != os.Stdout && b.writer != os.Stderr {
			if e := b.writer.Sync(); e != nil {
				err = e
				nBytes = 0
			}
		}
	}

	return nBytes, err

}

func (b *BaseHandler) filter(record LogRecord) bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	var isValidRecord int
	for _, filter := range b.filters {
		if filter.Filter(record) {
			isValidRecord += 1
		}
	}

	switch isValidRecord {
	case len(b.filters):
		return true
	default:
		return false
	}
}

func (b *BaseHandler) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.writer.Close()
}

// StreamHandler

type StreamHandler struct {
	bh *BaseHandler
}

func GetStreamHandler() (*StreamHandler, error) {
	return &StreamHandler{bh: &BaseHandler{writer: os.Stdout}}, nil
}

func (s *StreamHandler) SetFormatter(formatter Formatter) {
	s.bh.SetFormatter(formatter)
}

func (s *StreamHandler) emit(record LogRecord) (int, error) {
	return s.bh.emit(record)
}

func (s *StreamHandler) filter(record LogRecord) bool {
	return s.bh.filter(record)
}
func (s *StreamHandler) AddFilter(filter Filter) {
	s.bh.AddFilter(filter)
}

func (s *StreamHandler) SetLogLevel(level int) {
	s.bh.SetLogLevel(level)
}

func (s *StreamHandler) GetLogLevel() int {
	return s.bh.logLevel
}

func (s *StreamHandler) Close() {
	s.bh.Close()
}

//File Handler

type FileHandler struct {
	bh       *BaseHandler
	filename string
}

func GetFileHandler(filename string, flag int, perm os.FileMode) (*FileHandler, error) {
	file, err := os.OpenFile(filename, flag, perm)
	if err != nil {
		return nil, err
	}
	return &FileHandler{bh: &BaseHandler{writer: file}, filename: filename}, nil
}

func (f *FileHandler) SetLogLevel(level int) {
	f.bh.SetLogLevel(level)
}

func (f *FileHandler) GetLogLevel() int {
	return f.bh.logLevel
}

func (f *FileHandler) SetFormatter(formatter Formatter) {
	f.bh.SetFormatter(formatter)
}

func (f *FileHandler) emit(record LogRecord) (int, error) {
	return f.bh.emit(record)
}

func (f *FileHandler) filter(record LogRecord) bool {
	return f.bh.filter(record)
}
func (f *FileHandler) AddFilter(filter Filter) {
	f.bh.AddFilter(filter)
}

func (f *FileHandler) Close() {
	f.bh.Close()
}
