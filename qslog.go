package qslog

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

var (
	levelNames   = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	currentLevel = InfoLevel
	mu           sync.Mutex
	logger       *log.Logger
)

func init() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
}

func SetLevel(level Level) {
	mu.Lock()
	defer mu.Unlock()
	currentLevel = level
}

func GetLevel() Level {
	mu.Lock()
	defer mu.Unlock()
	return currentLevel
}

func SetOutput(w io.Writer) {
	mu.Lock()
	defer mu.Unlock()
	logger.SetOutput(w)
}

func SetPrefix(prefix string) {
	mu.Lock()
	defer mu.Unlock()
	logger.SetPrefix(prefix)
}

func SetFlags(flag int) {
	mu.Lock()
	defer mu.Unlock()
	logger.SetFlags(flag)
}

func shouldLog(level Level) bool {
	return level >= GetLevel()
}

func logMessage(level Level, msg string) {
	if !shouldLog(level) {
		return
	}

	levelStr := levelNames[level]

	logger.Printf("[%s] %s", levelStr, msg)
}

func Debug(v ...interface{}) {
	logMessage(DebugLevel, trimNewline(fmt.Sprintln(v...)))
}

func Info(v ...interface{}) {
	logMessage(InfoLevel, trimNewline(fmt.Sprintln(v...)))
}

func Warn(v ...interface{}) {
	logMessage(WarnLevel, trimNewline(fmt.Sprintln(v...)))
}

func Error(v ...interface{}) {
	logMessage(ErrorLevel, trimNewline(fmt.Sprintln(v...)))
}

func Fatal(v ...interface{}) {
	logMessage(FatalLevel, trimNewline(fmt.Sprintln(v...)))
	os.Exit(1)
}

func trimNewline(s string) string {
	if len(s) > 0 && s[len(s)-1] == '\n' {
		return s[:len(s)-1]
	}
	return s
}

func Debugf(format string, v ...interface{}) {
	logMessage(DebugLevel, fmt.Sprintf(format, v...))
}

func Infof(format string, v ...interface{}) {
	logMessage(InfoLevel, fmt.Sprintf(format, v...))
}

func Warnf(format string, v ...interface{}) {
	logMessage(WarnLevel, fmt.Sprintf(format, v...))
}

func Errorf(format string, v ...interface{}) {
	logMessage(ErrorLevel, fmt.Sprintf(format, v...))
}

func Fatalf(format string, v ...interface{}) {
	logMessage(FatalLevel, fmt.Sprintf(format, v...))
	os.Exit(1)
}
