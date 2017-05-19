// +build linux

package ConsoleColor

import (
	"fmt"
	"github.com/fatih/color"
)

func Print(c int, a ...interface{}) (int, error) {
	if GetOS() == OS_LINUX {
		defer locker.Unlock()
		locker.Lock()
		fc := GetForegroundColor(OS_LINUX, c)
		return color.New(color.Attribute(fc)).Print(a...)
	}
	return fmt.Print(a...)
}

func Printf(c int, format string, a ...interface{}) (int, error) {
	if GetOS() == OS_LINUX {
		defer locker.Unlock()
		locker.Lock()
		fc := GetForegroundColor(OS_LINUX, c)
		return color.New(color.Attribute(fc)).Printf(format, a...)
	}
	return fmt.Printf(format, a...)
}

func Println(c int, a ...interface{}) (int, error) {
	if GetOS() == OS_LINUX {
		defer locker.Unlock()
		locker.Lock()
		fc := GetForegroundColor(OS_LINUX, c)
		return color.New(color.Attribute(fc)).Println(a...)
	}
	return fmt.Println(a...)
}
