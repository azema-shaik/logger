package logger

import (
	"fmt"
	"strings"
)

//these are placeholders only

type StdFormatter struct {
	FormatString string
	DateFmt      string
}

func (f *StdFormatter) SetFormatter(formatString string, datefmt string) {
	f.FormatString = formatString
	f.DateFmt = datefmt
}

func (f *StdFormatter) GetFormatter() (string, string) {
	return f.FormatString, f.DateFmt
}

func (f *StdFormatter) Format(l LogRecord) string {
	formatString, datefmt := f.GetFormatter()

	var fmtString string
	switch formatString {
	case "":
		fmtString = "[%(asctime)s] : [%(levelname)s] : [%(funcName)s] : [%(msg)s]"
	default:
		fmtString = formatString
	}

	switch datefmt {
	case "":
		datefmt = "2006-01-02 15:04:05 PM"
	}
	sFileName := strings.Split(l.File, "/")

	logTime := l.Datetime.Format(datefmt)
	record := map[string]string{
		"%(asctime)s":   logTime,
		"%(funcName)s":  l.FuncName,
		"%(levelname)s": l.LevelName,
		"%(levelno)d":   fmt.Sprintf("%d", l.LevelNo),
		"%(lineno)d":    fmt.Sprintf("%d", l.LineNo),
		"%(name)s":      l.Name,
		"%(Lfilename)s": l.File,
		"%(msg)s":       l.Message,
		"%(Sfilename)s": sFileName[len(sFileName)-1],
	}

	for placeholder, replacement := range record {
		fmtString = strings.Replace(fmtString, placeholder, replacement, 1)
	}

	return fmtString + "\n"

}
