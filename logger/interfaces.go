package logger

type Handler interface {
	SetLogLevel(int)
	GetLogLevel() int
	emit(logRecord) (int, error)
	format(logRecord) string
	GetFilters() []Filter
	SetFormatter(Formatter)
	GetFormatter() Formatter
	Close()
}

type Formatter interface {
	Format(logRecord) string
}

type Filter interface {
	Filter(logRecord) bool
}
