package logging

import (
	"fmt"
	"runtime"

	"github.com/fatih/color"
)

var green func(a ...interface{}) string
var red func(a ...interface{}) string
var purple func(a ...interface{}) string
var yellow func(a ...interface{}) string
var cyan func(a ...interface{}) string
var blue func(a ...interface{}) string

func Goodln(v ...interface{}) {
	if runtime.GOOS == "linux" {
		green = color.New(color.FgGreen).SprintFunc()
		fmt.Print(green("[+] "))
		fmt.Println(v...)
	} else {
		fmt.Print("[+] ")
		for _, arg := range v {
			fmt.Print(arg)
		}
		fmt.Print("\n")
	}
}

func Badln(v ...interface{}) {
	if runtime.GOOS == "linux" {
		red = color.New(color.FgRed).SprintFunc()
		fmt.Print(red("[-] "))
		fmt.Println(v...)
	} else {
		fmt.Print("[-] ")
		fmt.Println(v...)
	}
}

func Infoln(v ...interface{}) {
	if runtime.GOOS == "linux" {
		purple = color.New(color.FgMagenta).SprintFunc()
		fmt.Print(purple("[*] "))
		fmt.Println(v...)
	} else {
		fmt.Print("[*] ")
		fmt.Println(v...)
	}
}
