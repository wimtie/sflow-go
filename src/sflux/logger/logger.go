package logger

import (
	"log"
	"fmt"
)

const (
    Loglevel_Error  = 0
    Loglevel_Warn   = 1
    Loglevel_Info   = 2
    Loglevel_Debug  = 3
)

type Log struct {
	logLevel int
}

func NewLog(level int) (Log) {
	return Log { level }
}

func (l Log) Error(e error) {
	log.Fatal(e)
}

func (l Log) Warn(msg string, args ...interface{}) {
	l.log(msg, Loglevel_Warn, args...)
}

func (l Log) Info(msg string, args ...interface{}) {
	l.log(msg, Loglevel_Info, args...)
}

func (l Log) Debug(msg string, args ...interface{}) {
	l.log(msg, Loglevel_Debug, args...)
}

func (l Log) log(msg string, level int, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	if (l.logLevel >= level) {
		log.Println(msg)
	}
}
