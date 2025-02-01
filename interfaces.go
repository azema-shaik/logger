package main

type Handler interface {
	SetLogLevel(int)
	GetLogLevel() int
	emit(logRecord) (int, error)
	format(logRecord) string
	GetFilters() []Filter
	Close()
}

type Formatter interface {
	Format(logRecord) string
}

type Filter interface {
	Filter(logRecord) bool
}
