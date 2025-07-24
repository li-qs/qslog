package qslog

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

var levelNames = []string{"DEBUG", "INFO", "WARN", "ERROR"}

var currentLevel = InfoLevel
var mu sync.Mutex

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

func shouldLog(level Level) bool {
	return level >= GetLevel()
}

func logMessage(level Level, msg string) {
	if !shouldLog(level) {
		return
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	levelStr := levelNames[level]

	log.Printf("[%s] [%s] %s", now, levelStr, msg)
}

func Debug(v ...interface{}) {
	logMessage(DebugLevel, fmt.Sprint(v...))
}

func Info(v ...interface{}) {
	logMessage(InfoLevel, fmt.Sprint(v...))
}

func Warn(v ...interface{}) {
	logMessage(WarnLevel, fmt.Sprint(v...))
}

func Error(v ...interface{}) {
	logMessage(ErrorLevel, fmt.Sprint(v...))
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

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
}
