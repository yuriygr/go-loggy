package loggy

import (
	"fmt"
	"time"
)

var (
	black   = color("\033[1;30m%s\033[0m")
	red     = color("\033[1;31m%s\033[0m")
	green   = color("\033[1;32m%s\033[0m")
	yellow  = color("\033[1;33m%s\033[0m")
	purple  = color("\033[1;34m%s\033[0m")
	magenta = color("\033[1;35m%s\033[0m")
	teal    = color("\033[1;36m%s\033[0m")
	white   = color("\033[1;37m%s\033[0m")

	success = color("\033[1;32m%s\033[0m")
	info    = color("\033[1;34m%s\033[0m")
	notice  = color("\033[1;36m%s\033[0m")
	warning = color("\033[1;33m%s\033[0m")
	error   = color("\033[1;31m%s\033[0m")
	debug   = color("\033[0;36m%s\033[0m")
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
	l.Writeln(success(message))
}

// Info - Запись в лог
func (l *Logger) Info(message interface{}) {
	l.Writeln(info(message))
}

// Notice - Запись в лог
func (l *Logger) Notice(message interface{}) {
	l.Writeln(notice(message))
}

// Warning - Запись в лог
func (l *Logger) Warning(message interface{}) {
	l.Writeln(warning(message))
}

// Error - Запись в лог
func (l *Logger) Error(message interface{}) {
	l.Writeln(error(message))
}

// Debug - Запись в лог
func (l *Logger) Debug(message interface{}) {
	l.Writeln(debug(message))
}
