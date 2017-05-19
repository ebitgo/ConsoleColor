// +build windows

package ConsoleColor

import (
	"fmt"
)

func Print(c int, a ...interface{}) (int, error) {
	if GetOS() == OS_WINDOWS {
		defer locker.Unlock()
		defer ResetColor()
		locker.Lock()
		fc := GetForegroundColor(OS_WINDOWS, c)
		Textbackground(fc)
	}
	return fmt.Print(a...)
}

func Printf(c int, format string, a ...interface{}) (int, error) {
	if GetOS() == OS_WINDOWS {
		defer locker.Unlock()
		defer ResetColor()
		locker.Lock()
		fc := GetForegroundColor(OS_WINDOWS, c)
		Textbackground(fc)
	}
	return fmt.Printf(format, a...)
}

func Println(c int, a ...interface{}) (int, error) {
	if GetOS() == OS_WINDOWS {
		defer locker.Unlock()
		defer ResetColor()
		locker.Lock()
		fc := GetForegroundColor(OS_WINDOWS, c)
		Textbackground(fc)
	}
	return fmt.Println(a...)
}

func PrintXY(c int, x, y int16, a ...interface{}) (int, error) {
	if GetOS() == OS_WINDOWS {
		defer locker.Unlock()
		defer ResetColor()
		locker.Lock()
		fc := GetForegroundColor(OS_WINDOWS, c)
		Textbackground(fc)
		Gotoxy(x, y)
	}
	return fmt.Print(a...)
}

func PrintfXY(c int, x, y int16, format string, a ...interface{}) (int, error) {
	if GetOS() == OS_WINDOWS {
		defer locker.Unlock()
		defer ResetColor()
		locker.Lock()
		fc := GetForegroundColor(OS_WINDOWS, c)
		Textbackground(fc)
		Gotoxy(x, y)
	}
	return fmt.Printf(format, a...)
}

func PrintlnXY(c int, x, y int16, a ...interface{}) (int, error) {
	if GetOS() == OS_WINDOWS {
		defer locker.Unlock()
		defer ResetColor()
		locker.Lock()
		fc := GetForegroundColor(OS_WINDOWS, c)
		Textbackground(fc)
		Gotoxy(x, y)
	}
	return fmt.Println(a...)
}
