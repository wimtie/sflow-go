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

func (l Log) Warn(msg string) {
	l.log(msg, Loglevel_Warn)
}

func (l Log) Info(msg string) {
	l.log(msg, Loglevel_Info)
}

func (l Log) Debug(msg string, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	l.log(msg, Loglevel_Debug)
}

func (l Log) log(msg string, level int) {
	if (l.logLevel >= level) {
		log.Println(msg)
	}
}
