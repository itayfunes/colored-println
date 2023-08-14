package main

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	red   = color.New(color.FgRed).SprintFunc()
	green = color.New(color.FgGreen).SprintFunc()
	blue  = color.New(color.FgBlue).SprintFunc()
)

func logWithPrefix(prefix string, colorFunc func(...interface{}) string, message string) {
	fmt.Printf("[%s] %s\n", colorFunc(prefix), message)
}

func badln(message string) {
	logWithPrefix("-", red, message)
}

func goodln(message string) {
	logWithPrefix("+", green, message)
}

func infoln(message string) {
	logWithPrefix("*", blue, message)
}

func main() {
	badln("This is a bad message.")
	goodln("This is a good message.")
	infoln("This is an info message.")
}
