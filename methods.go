package loggy

import (
	"fmt"
	"time"
)

// Writeln - Запись в лог
func (l *Logger) Writeln(message interface{}) {
	currentTime := time.Now()
	fmt.Fprintf(l.config.Writer, "%s - %s\n", currentTime.Format(l.config.TimeFormat), message)
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
