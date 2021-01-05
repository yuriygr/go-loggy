# Humble & Colorful logging library for Golang
[![Go Report Card](https://goreportcard.com/badge/github.com/yuriygr/go-loggy)](https://goreportcard.com/report/github.com/yuriygr/go-loggy)
[![Build Status](https://travis-ci.org/yuriygr/go-loggy.svg?branch=master)](https://travis-ci.org/yuriygr/go-loggy)
[![](https://godoc.org/github.com/yuriygr/go-loggy?status.svg)](https://pkg.go.dev/github.com/yuriygr/go-loggy?tab=doc)

## About

Все просто, сторонние решения слишком сложные и тяжелые, а это... ну, мои задачи оно решает.

## Install

`go get -u github.com/yuriygr/go-loggy`

## Features

* Humble
* Colorful
* No dependencies

## Quick Start

```go
config := loggy.LoggerConfig{
    Writer:     os.Stderr,
    TimeFormat: "2006-01-02 15:04:05",
}
logger := loggy.NewLogger(config)
```

## Examples

```go
logger.Success("Success text")

logger.Info("Info text")

logger.Notice("Notice text")

logger.Warning("Warning text")

logger.Error("Error text")

logger.Debug("Debug text")
```
