package logger

type Handler interface {
	SetLogLevel(int)
	GetLogLevel() int
	emit(LogRecord) (int, error)
	format(LogRecord) string
	GetFilters() []Filter
	SetFormatter(Formatter)
	GetFormatter() Formatter
	Close()
}

type Formatter interface {
	SetFormatter(formatString string, datefmt string)
	GetFormatter() (string, string)
	Format(LogRecord) string
}

type Filter interface {
	Filter(LogRecord) bool
}
