package main

import (
	"log"
	"os"
)

type ExtendedLog struct {
	*log.Logger
	logLevel LogLevel
}

type LogLevel int

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
)

func (l LogLevel) IsValid() bool {
	switch l {
	case LogLevelError, LogLevelWarning, LogLevelInfo:
		return true
	default:
		return false
	}

}

func NewExtendedLogger() *ExtendedLog {
	return &ExtendedLog{
		Logger:   log.New(os.Stderr, "", log.LstdFlags),
		logLevel: LogLevelError,
	}

}

func (l *ExtendedLog) SetLogLevel(logLvl LogLevel) {
	if !logLvl.IsValid() {
		return
	}
	l.logLevel = logLvl
}

func (l *ExtendedLog) Infoln(msg string) {
	l.println(LogLevelInfo, "INFO ", msg)
}

func (l *ExtendedLog) Warnln(msg string) {
	l.println(LogLevelWarning, "WARN ", msg)
}

func (l *ExtendedLog) Errorln(msg string) {
	l.println(LogLevelError, "ERR ", msg)
}

func (l *ExtendedLog) println(srcLogLvl LogLevel, prefix, msg string) {
	if l.logLevel < srcLogLvl {
		return
	}

	l.Logger.Println(prefix + msg)
}
