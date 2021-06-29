package logmod

type logger struct {
}

type Logger interface {
	Info(string)
	Debug(string)
	Error(string)
}

func NewLogger() Logger {
	return &logger{}
}

func (l *logger) Info(s string) {

}

func (l *logger) Debug(s string) {

}

func (l *logger) Error(s string) {

}
