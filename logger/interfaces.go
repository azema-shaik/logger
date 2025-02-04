package logger

type Handler interface {
	SetLogLevel(int)
	emit(LogRecord) (int, error)
	SetFormatter(Formatter)
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

type LoggerLike interface{}
