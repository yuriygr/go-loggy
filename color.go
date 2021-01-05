package loggy

import "fmt"

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
