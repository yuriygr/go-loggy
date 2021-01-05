package loggy

import (
	"io"
	"os"
)

// NewLogger - Создаем экземпляр рассыльщика писем
func NewLogger(config LoggerConfig) *Logger {
	defaultConfig := LoggerConfig{
		Writer:     os.Stderr,
		TimeFormat: "2006-01-02 15:04:05",
	}

	if config.Writer != nil {
		defaultConfig.Writer = config.Writer
	}

	if config.TimeFormat != "" {
		defaultConfig.TimeFormat = config.TimeFormat
	}

	return &Logger{defaultConfig}
}

// Logger - Структура логгера
type Logger struct {
	config LoggerConfig
}

// LoggerConfig - Структура настроек
type LoggerConfig struct {
	Writer     io.Writer
	TimeFormat string
}
