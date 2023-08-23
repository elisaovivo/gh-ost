/*
   Copyright 2022 GitHub Inc.
         See https://github.com/github/gh-ost/blob/master/LICENSE
*/

package base

import (
	"bytes"
	"io"

	"github.com/openark/golib/log"
)

type simpleLogger struct {
	buf *bytes.Buffer
}

func NewDefaultLogger() *simpleLogger {
	return &simpleLogger{
		buf: bytes.NewBufferString(""),
	}
}

func (l *simpleLogger) Debug(args ...interface{}) {

	s := log.Debug(args[0].(string), args[1:])
	l.buf.WriteString(s)
	l.buf.WriteString("\n")
}

func (l *simpleLogger) Debugf(format string, args ...interface{}) {
	s := log.Debugf(format, args...)
	l.buf.WriteString(s)
	l.buf.WriteString("\n")
}

func (l *simpleLogger) Info(args ...interface{}) {
	s := log.Info(args[0].(string), args[1:])
	l.buf.WriteString(s)
	l.buf.WriteString("\n")
}

func (l *simpleLogger) Infof(format string, args ...interface{}) {
	s := log.Infof(format, args...)
	l.buf.WriteString(s)
	l.buf.WriteString("\n")
}

func (l *simpleLogger) Warning(args ...interface{}) error {
	err := log.Warning(args[0].(string), args[1:])
	l.buf.WriteString(err.Error())
	l.buf.WriteString("\n")
	return err
}

func (l *simpleLogger) Warningf(format string, args ...interface{}) error {
	err := log.Warningf(format, args...)
	l.buf.WriteString(err.Error())
	l.buf.WriteString("\n")
	return err
}

func (l *simpleLogger) Error(args ...interface{}) error {
	err := log.Error(args[0].(string), args[1:])
	l.buf.WriteString(err.Error())
	l.buf.WriteString("\n")
	return err
}

func (l *simpleLogger) Errorf(format string, args ...interface{}) error {
	err := log.Errorf(format, args...)
	l.buf.WriteString(err.Error())
	l.buf.WriteString("\n")
	return err
}

func (l *simpleLogger) Errore(err error) error {
	e := log.Errore(err)
	l.buf.WriteString(e.Error())
	l.buf.WriteString("\n")
	return e
}

func (l *simpleLogger) Fatal(args ...interface{}) error {
	e := log.Fatal(args[0].(string), args[1:])
	l.buf.WriteString(e.Error())
	l.buf.WriteString("\n")
	return e
}

func (l *simpleLogger) Fatalf(format string, args ...interface{}) error {
	e := log.Fatalf(format, args...)
	l.buf.WriteString(e.Error())
	l.buf.WriteString("\n")
	return e
}

func (l *simpleLogger) Fatale(err error) error {
	e := log.Fatale(err)
	l.buf.WriteString(e.Error())
	l.buf.WriteString("\n")
	return e
}

func (*simpleLogger) SetLevel(level log.LogLevel) {
	log.SetLevel(level)
}

func (*simpleLogger) SetPrintStackTrace(printStackTraceFlag bool) {
	log.SetPrintStackTrace(printStackTraceFlag)
}

func (l *simpleLogger) GetBuf() io.Reader {
	return l.buf
}
