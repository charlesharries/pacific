package main

import "log"

type Logger struct{}

func (l Logger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (l Logger) Verbose() bool {
	return true
}
