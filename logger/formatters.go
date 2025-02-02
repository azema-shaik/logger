package logger

import (
	"fmt"
	"strings"
)

//these are placeholders only

type StdFormatter struct {
	formatString string
	datefmt      string
}

func (f *StdFormatter) Format(l logRecord) string {
	var datefmt string
	switch f.datefmt {
	case "":
		datefmt = "2006-01-02 15:04:05 PM"
	default:
		datefmt = f.datefmt

	}
	var fmtString string
	switch fmtString {
	case "":
		fmtString = "[%(asctime)s] : [%(levelname)s] : [%(funcName)s] : [%(msg)s]"
	default:
		fmtString = f.formatString
	}
	logTime := l.Datetime.Format(datefmt)
	record := map[string]string{
		"%(asctime)s":   logTime,
		"%(funcName)s":  l.FuncName,
		"%(levelname)s": l.LevelName,
		"%(levelno)d":   fmt.Sprintf("%d", l.LevelNo),
		"%(lineno)d":    fmt.Sprintf("%d", l.LineNo),
		"%(name)s":      l.Name,
		"%(filename)s":  l.File,
		"%(msg)s":       l.Message,
	}

	for placeholder, replacement := range record {
		fmtString = strings.Replace(fmtString, placeholder, replacement, 1)
	}

	return fmtString + "\n"
}
