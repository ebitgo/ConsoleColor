package ConsoleColor

import (
	"errors"
	"github.com/fatih/color"
	"sync"
)

const (
	FOREGROUND = iota
	BACKGROUND

	C_BLACK = iota
	C_RED
	C_GREEN
	C_YELLOW
	C_BLUE
	C_MAGENTA
	C_CYAN
	C_WHITE

	// Windows console 颜色定义
	FOREGROUND_BLACK       = 0x00 // black.
	FOREGROUND_DARKBLUE    = 0x01 // dark blue.
	FOREGROUND_DARKGREEN   = 0x02 // dark green.
	FOREGROUND_DARKSKYBLUE = 0x03 // dark skyblue.
	FOREGROUND_DARKRED     = 0x04 // dark red.
	FOREGROUND_DARKPINK    = 0x05 // dark pink.
	FOREGROUND_DARKYELLOW  = 0x06 // dark yellow.
	FOREGROUND_DARKWHITE   = 0x07 // dark white.
	FOREGROUND_DARKGRAY    = 0x08 // dark gray.
	FOREGROUND_BLUE        = 0x09 // blue.
	FOREGROUND_GREEN       = 0x0a // green.
	FOREGROUND_SKYBLUE     = 0x0b // skyblue.
	FOREGROUND_RED         = 0x0c // red.
	FOREGROUND_PINK        = 0x0d // pink.
	FOREGROUND_YELLOW      = 0x0e // yellow.
	FOREGROUND_WHITE       = 0x0f // white.

	BACKGROUND_DARKBLUE    = 0x10 // dark blue.
	BACKGROUND_DARKGREEN   = 0x20 // dark green.
	BACKGROUND_DARKSKYBLUE = 0x30 // dark skyblue.
	BACKGROUND_DARKRED     = 0x40 // dark red.
	BACKGROUND_DARKPINK    = 0x50 // dark pink.
	BACKGROUND_DARKYELLOW  = 0x60 // dark yellow.
	BACKGROUND_DARKWHITE   = 0x70 // dark white.
	BACKGROUND_DARKGRAY    = 0x80 // dark gray.
	BACKGROUND_BLUE        = 0x90 // blue.
	BACKGROUND_GREEN       = 0xa0 // green.
	BACKGROUND_SKYBLUE     = 0xb0 // skyblue.
	BACKGROUND_RED         = 0xc0 // red.
	BACKGROUND_PINK        = 0xd0 // pink.
	BACKGROUND_YELLOW      = 0xe0 // yellow.
	BACKGROUND_WHITE       = 0xf0 // white.
)

var locker *sync.Mutex = new(sync.Mutex)

func GetForegroundColor(os int, c int) int {
	if os == OS_WINDOWS {
		switch c {
		case C_BLACK:
			return FOREGROUND_BLACK
		case C_RED:
			return FOREGROUND_RED
		case C_GREEN:
			return FOREGROUND_DARKGREEN
		case C_YELLOW:
			return FOREGROUND_YELLOW
		case C_BLUE:
			return FOREGROUND_BLUE
		case C_MAGENTA:
			return FOREGROUND_DARKRED
		case C_CYAN:
			return FOREGROUND_GREEN
		case C_WHITE:
			return FOREGROUND_WHITE
		}
	} else {
		switch c {
		case C_BLACK:
			return int(color.FgBlack)
		case C_RED:
			return int(color.FgRed)
		case C_GREEN:
			return int(color.FgGreen)
		case C_YELLOW:
			return int(color.FgYellow)
		case C_BLUE:
			return int(color.FgBlue)
		case C_MAGENTA:
			return int(color.FgMagenta)
		case C_CYAN:
			return int(color.FgCyan)
		case C_WHITE:
			return int(color.FgWhite)
		}
	}
	panic(errors.New("Undefined foreground color!"))
}

func GetBackgroundColor(os int, c int) int {
	if os == OS_WINDOWS {
		switch c {
		case C_BLACK:
			return FOREGROUND_BLACK
		case C_RED:
			return BACKGROUND_RED
		case C_GREEN:
			return BACKGROUND_GREEN
		case C_YELLOW:
			return BACKGROUND_YELLOW
		case C_BLUE:
			return BACKGROUND_BLUE
		case C_MAGENTA:
			return BACKGROUND_DARKRED
		case C_CYAN:
			return BACKGROUND_DARKSKYBLUE
		case C_WHITE:
			return BACKGROUND_WHITE
		}
	} else {
		switch c {
		case C_BLACK:
			return int(color.BgBlack)
		case C_RED:
			return int(color.BgRed)
		case C_GREEN:
			return int(color.BgGreen)
		case C_YELLOW:
			return int(color.BgYellow)
		case C_BLUE:
			return int(color.BgBlue)
		case C_MAGENTA:
			return int(color.BgMagenta)
		case C_CYAN:
			return int(color.BgCyan)
		case C_WHITE:
			return int(color.BgWhite)
		}
	}
	panic(errors.New("Undefined background color!"))
}
