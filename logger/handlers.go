package logger

import (
	"os"
)

func setDefaultFormatter(handler Handler) {
	if handler.GetFormatter() == nil {
		handler.SetFormatter(&StdFormatter{})
	}
}

// StreamHandler

type StreamHandler struct {
	bh *BaseHandler
}

func GetStreamHandler() *StreamHandler {
	return &StreamHandler{bh: &BaseHandler{writer: os.Stdout}}
}

//File Handler

type FileHandler struct {
	bh *BaseHandler
}

func GetFileHandler(filename string, flag int, perm os.FileMode) *FileHandler {
	file, _ := os.OpenFile(filename, flag, perm)
	return &FileHandler{bh: &BaseHandler{writer: file}}
}
