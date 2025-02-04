package logger

type rootLogger struct {
	logger Logger
}

func (r *rootLogger) SetLevel(level int) {
	r.logger.Level = level
}

func (r *rootLogger) AddHandlers(handler Handler) {
	r.logger.Handlers = append(r.logger.Handlers, handler)
}

func (r *rootLogger) log(message string, level int) {
	r.logger.log(message, level)
}

func (r *rootLogger) Debug(message string) {
	r.log(message, DEBUG)
}

func (r *rootLogger) Info(message string) {
	r.log(message, INFO)
}

func (r *rootLogger) Warning(message string) {
	r.log(message, WARNING)
}
func (r *rootLogger) Critical(message string) {
	r.log(message, CRITICAL)
}

func (r *rootLogger) Close() {
	for _, hdlr := range r.logger.Handlers {
		hdlr.Close()
	}
}

var root *rootLogger
