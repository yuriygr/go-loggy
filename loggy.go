package loggy

import (
	"fmt"
	"time"
)

var (
	Black   = color("\033[1;30m%s\033[0m")
	Red     = color("\033[1;31m%s\033[0m")
	Green   = color("\033[1;32m%s\033[0m")
	Yellow  = color("\033[1;33m%s\033[0m")
	Purple  = color("\033[1;34m%s\033[0m")
	Magenta = color("\033[1;35m%s\033[0m")
	Teal    = color("\033[1;36m%s\033[0m")
	White   = color("\033[1;37m%s\033[0m")

	Success = color("\033[1;32m%s\033[0m")
	Info    = color("\033[1;34m%s\033[0m")
	Notice  = color("\033[1;36m%s\033[0m")
	Warning = color("\033[1;33m%s\033[0m")
	Error   = color("\033[1;31m%s\033[0m")
	Debug   = color("\033[0;36m%s\033[0m")
)

func color(colorString string) func(...interface{}) string {
	return func(args ...interface{}) string {
		return fmt.Sprintf(colorString, fmt.Sprint(args...))
	}
}

// NewLogger - Создаем экземпляр рассыльщика писем
func NewLogger() *Logger {
	config := LoggerConfig{
		TimeFormat: "2006-01-02 15:04:05",
	}
	return &Logger{config}
}

// Logger - Структура логгера
type Logger struct {
	config LoggerConfig
}

// LoggerConfig - Структура настроек
type LoggerConfig struct {
	TimeFormat string
}

// Writeln - Запись в лог
func (l *Logger) Writeln(message interface{}) {
	currentTime := time.Now()
	fmt.Printf("%s - %s", currentTime.Format(l.config.TimeFormat), message)
	fmt.Println()
}

// Success - Запись в лог
func (l *Logger) Success(message interface{}) {
	l.Writeln(Success(message))
}

// Info - Запись в лог
func (l *Logger) Info(message interface{}) {
	l.Writeln(Info(message))
}

// Notice - Запись в лог
func (l *Logger) Notice(message interface{}) {
	l.Writeln(Notice(message))
}

// Warning - Запись в лог
func (l *Logger) Warning(message interface{}) {
	l.Writeln(Warning(message))
}

// Error - Запись в лог
func (l *Logger) Error(message interface{}) {
	l.Writeln(Error(message))
}
