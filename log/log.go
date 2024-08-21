package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

const(
	DEBUG = iota
	INFO
	WARNING
	ERROR
	FATAL
)

type Level int

func(l Level) String() string{
	if l >=DEBUG && l <= FATAL{
		return LevelMap[l]
	}
	return fmt.Sprintf("[Level(%d)]", l)
}

var(
	defaultFlags = log.LstdFlags | log.Lmicroseconds | log.Lshortfile
	defaultPrefix = ""
	defaultLevel = WARNING
	LevelMap = []string{
		"[DEBUG]: ",
		"[INFO]:",
		"[WARNING]: ",
		"[ERROR]: ",
		"[FATAL]: ",
	}
)

type Logger interface{
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	SetLevel(str string)
	SetOutput(out io.Writer)
	SetPrefix(prefix string)
	SetFlags(flags int)
}

type _logger struct{
	originLog *log.Logger
	level int
	output io.Writer
}

var _ Logger = (*_logger)(nil)

func (l *_logger) Level() string {
	return LevelMap[l.level]
}

func (l *_logger) Info(v ...interface{}){
	l.logf(INFO, nil, v ...)
}

func (l *_logger) Debug(v ...interface{}){
	l.logf(DEBUG, nil, v ...)
}

func (l *_logger) Warn(v ... interface{}){
	l.logf(WARNING, nil, v ...)
}

func (l *_logger) Error(v ... interface{}){
	l.logf(ERROR, nil, v ...)
}

func (l *_logger) Fatal(v ... interface{}){
	l.logf(FATAL, nil, v ...)
}

func (l *_logger) Infof(format string, v ...interface{}){
	l.logf(INFO, &format, v ...)
}

func (l *_logger) Debugf(format string, v ...interface{}){
	l.logf(DEBUG, &format, v ...)
}

func (l *_logger) Warnf(format string, v ... interface{}){
	l.logf(WARNING, &format, v ...)
}

func (l *_logger) Errorf(format string, v ... interface{}){
	l.logf(ERROR, &format, v ...)
}

func (l *_logger) Fatalf(format string, v ... interface{}){
	l.logf(FATAL, &format, v ...)
}


func (l *_logger) logf(lv Level, format *string, v ...interface{}){
	if int(lv) < l.level{
		return 
	}
	msg := lv.String()
	if format != nil{
		msg += fmt.Sprintf(*format, v...)
	}else{
		msg += fmt.Sprint(v...)
	}
	_ = l.originLog.Output(4, msg)
	if lv ==FATAL{
		os.Exit(1)
	}
}


func (l *_logger) SetLevel(level string){
	switch strings.ToLower(level){
	case "debug":
		l.level = DEBUG
		break
	case "info":
		l.level = INFO
		break
	case "warning":
		l.level = WARNING
		break
	case "warn":
		l.level = WARNING
		break
	case "error":
		l.level = ERROR
		break
	case "fatal":
		l.level = FATAL
		break
	default:
		l.level = WARNING
	}
}

func (l *_logger) SetOutput(output io.Writer){
	l.originLog.SetOutput(output)
} 

func (l *_logger) SetPrefix(prefix string){
	l.originLog.SetPrefix(prefix)
}

func (l *_logger) SetFlags(flag int){
	l.originLog.SetFlags(flag)
}

var logger Logger = &_logger{
	originLog: log.New(os.Stdout, defaultPrefix, defaultFlags),
	level: defaultLevel,
	output: os.Stdout,
}

func SetOutput(w io.Writer){
	logger.SetOutput(w)
}

func SetLevel(level string){
	logger.SetLevel(level)
}

func SetPrefix(prefix string){
	logger.SetPrefix(prefix)
}

func SetFlags(flags int){
	logger.SetFlags(flags)
}

func Debug(v ...interface{}){
	logger.Debug(v...)
}

func Info(v ...interface{}){
	logger.Info(v...)
}

func Warn(v ...interface{}){
	logger.Warn(v...)
}

func Error(v ...interface{}){
	logger.Error(v...)
}
func Fatal(v ...interface{}){
	logger.Fatal(v...)
}


func Debugf(format string, v ...interface{}){
	logger.Debugf(format, v...)
}

func Infof(format string, v ...interface{}){
	logger.Infof(format, v...)
}

func Warnf(format string, v ...interface{}){
	logger.Warnf(format, v...)
}

func Errorf(format string, v ...interface{}){
	logger.Errorf(format, v...)
}
func Fatalf(format string, v ...interface{}){
	logger.Fatalf(format, v...)
}

