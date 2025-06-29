package main

import (
	"fmt"
	"os"
	"time"
)

type logger interface {
	Info(string)
}

type screenLogger struct{}

func (s screenLogger) Info(message string) {
	fmt.Printf("%v -Info:%s", time.Now(), message)
}

type Filelogger struct {
	file *os.File
}

func (l *Filelogger) Info(message string) {
	fmt.Fprintf(l.file, "%v-Info %s", time.Now(), message)
}
func NewFilelogger(filename string) *Filelogger {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)	
	}
	return &Filelogger{file}
}

func main() {
	var log logger
	log = NewFilelogger("Log.txt")
	log.Info("Hello Rajesh!")
}
