package logmod

import "log"

type logger struct {
}

type Logger interface {
	Info(string)
	Debug(string)
	Error(string)
}

func NewLogger() Logger {

	log.Println("Create the new logger")
	return &logger{}
}

func (l *logger) Info(s string) {
	log.Println("Writing the Info log ")

}

func (l *logger) Debug(s string) {
	log.Println("Writing the Debug log ")

}

func (l *logger) Error(s string) {
	log.Println("Writing the Error log ")

}
