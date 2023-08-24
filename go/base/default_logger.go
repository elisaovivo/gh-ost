/*
   Copyright 2022 GitHub Inc.
         See https://github.com/github/gh-ost/blob/master/LICENSE
*/

package base

import (
	"github.com/openark/golib/log"
)

type simpleLogger struct {
	buf chan string
}

func NewDefaultLogger() *simpleLogger {
	return &simpleLogger{
		buf: make(chan string),
	}
}

func (l *simpleLogger) Debug(args ...interface{}) {

	log.Debug(args[0].(string), args[1:])

}

func (l *simpleLogger) Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func (l *simpleLogger) Info(args ...interface{}) {
	s := log.Info(args[0].(string), args[1:])
	l.buf <- s
}

func (l *simpleLogger) Infof(format string, args ...interface{}) {
	s := log.Infof(format, args...)
	l.buf <- s
}

func (l *simpleLogger) Warning(args ...interface{}) error {
	err := log.Warning(args[0].(string), args[1:])
	if err != nil {
		l.buf <- err.Error()
	}

	return err
}

func (l *simpleLogger) Warningf(format string, args ...interface{}) error {
	err := log.Warningf(format, args...)
	if err != nil {
		l.buf <- err.Error()
	}
	return err
}

func (l *simpleLogger) Error(args ...interface{}) error {
	err := log.Error(args[0].(string), args[1:])
	if err != nil {
		l.buf <- err.Error()
	}
	return err
}

func (l *simpleLogger) Errorf(format string, args ...interface{}) error {
	err := log.Errorf(format, args...)
	if err != nil {
		l.buf <- err.Error()
	}

	return err
}

func (l *simpleLogger) Errore(err error) error {
	e := log.Errore(err)
	if err != nil {
		l.buf <- err.Error()
	}

	return e
}

func (l *simpleLogger) Fatal(args ...interface{}) error {
	e := log.Fatal(args[0].(string), args[1:])
	if e != nil {
		l.buf <- e.Error()
	}

	return e
}

func (l *simpleLogger) Fatalf(format string, args ...interface{}) error {
	e := log.Fatalf(format, args...)
	if e != nil {
		l.buf <- e.Error()
	}
	return e
}

func (l *simpleLogger) Fatale(err error) error {
	e := log.Fatale(err)
	if e != nil {
		l.buf <- e.Error()
	}
	return e
}

func (*simpleLogger) SetLevel(level log.LogLevel) {
	log.SetLevel(level)
}

func (*simpleLogger) SetPrintStackTrace(printStackTraceFlag bool) {
	log.SetPrintStackTrace(printStackTraceFlag)
}

func (l *simpleLogger) GetBuf() chan string {
	return l.buf
}
