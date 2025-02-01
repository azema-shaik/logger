package main

type Handler interface {
	Emit(logRecord) (int, error)
	format(logRecord) string
}

type Formatter interface {
	Format(logRecord) string
}

type Filter interface {
	Filter(logRecord) bool
}
