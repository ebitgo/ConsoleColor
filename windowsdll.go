// +build windows

package ConsoleColor

import (
	"syscall"
	"unsafe"
)

type (
	HANDLE uintptr
	WORD   uint16
	DWORD  uint32
	TCHAR  rune
)

const (
	STD_OUTPUT_HANDLE = 0xFFFFFFF5
)

type COORD struct {
	X, Y int16
}

type SMALL_RECT struct {
	Left, Top, Right, Bottom int16
}

type CONSOLE_SCREEN_BUFFER_INFO struct {
	DwSize              COORD
	DwCursorPosition    COORD
	WAttributes         WORD
	SrWindow            SMALL_RECT
	DwMaximumWindowSize COORD
}

var OrigSrWindowColor WORD

var (
	modkernel32                    = syscall.NewLazyDLL("kernel32.dll")
	procFillConsoleOutputAttribute = modkernel32.NewProc("FillConsoleOutputAttribute")
	procFillConsoleOutputCharacter = modkernel32.NewProc("FillConsoleOutputCharacterW")
	procGetStdHandle               = modkernel32.NewProc("GetStdHandle")
	procGetConsoleScreenBufferInfo = modkernel32.NewProc("GetConsoleScreenBufferInfo")
	procSetConsoleCursorPosition   = modkernel32.NewProc("SetConsoleCursorPosition")
	procSetConsoleTextAttribute    = modkernel32.NewProc("SetConsoleTextAttribute")
)

func init() {
	if GetOS() == OS_WINDOWS {
		hConsole := GetStdHandle(STD_OUTPUT_HANDLE)
		var csbi *CONSOLE_SCREEN_BUFFER_INFO
		csbi = GetConsoleScreenBufferInfo(hConsole)
		OrigSrWindowColor = csbi.WAttributes
	}
}

func COORD2DWORD(c COORD) DWORD {
	return DWORD(int32(c.Y)<<16 + int32(c.X))
}

func FillConsoleOutputAttribute(hConsoleOutput HANDLE, wAttribute WORD, nLength DWORD, dwWriteCoord COORD) *DWORD {
	var lpNumberOfAttrsWritten DWORD
	ret, _, _ := procFillConsoleOutputAttribute.Call(
		uintptr(hConsoleOutput),
		uintptr(wAttribute),
		uintptr(nLength),
		uintptr(COORD2DWORD(dwWriteCoord)),
		uintptr(unsafe.Pointer(&lpNumberOfAttrsWritten)))
	if ret == 0 {
		return nil
	}
	return &lpNumberOfAttrsWritten
}

func FillConsoleOutputCharacter(hConsoleOutput HANDLE, cCharacter TCHAR, nLength DWORD, dwWriteCoord COORD) *DWORD {
	var lpNumberOfAttrsWritten DWORD
	ret, _, _ := procFillConsoleOutputCharacter.Call(
		uintptr(hConsoleOutput),
		uintptr(cCharacter),
		uintptr(nLength),
		uintptr(COORD2DWORD(dwWriteCoord)),
		uintptr(unsafe.Pointer(&lpNumberOfAttrsWritten)))
	if ret == 0 {
		return nil
	}
	return &lpNumberOfAttrsWritten
}

func GetStdHandle(nStdHandle DWORD) HANDLE {
	ret, _, _ := procGetStdHandle.Call(uintptr(nStdHandle))
	return HANDLE(ret)
}

func GetConsoleScreenBufferInfo(hConsoleOutput HANDLE) *CONSOLE_SCREEN_BUFFER_INFO {
	var csbi CONSOLE_SCREEN_BUFFER_INFO
	ret, _, _ := procGetConsoleScreenBufferInfo.Call(
		uintptr(hConsoleOutput),
		uintptr(unsafe.Pointer(&csbi)))
	if ret == 0 {
		return nil
	}
	return &csbi
}

func SetConsoleCursorPosition(hConsoleOutput HANDLE, dwCursorPosition COORD) bool {
	ret, _, _ := procSetConsoleCursorPosition.Call(
		uintptr(hConsoleOutput),
		uintptr(COORD2DWORD(dwCursorPosition)))
	return ret != 0
}

func SetConsoleTextAttribute(hConsoleOutput HANDLE, wAttributes WORD) bool {
	ret, _, _ := procSetConsoleTextAttribute.Call(
		uintptr(hConsoleOutput),
		uintptr(wAttributes))
	return ret != 0
}

func ClrSrc() {
	hConsole := GetStdHandle(STD_OUTPUT_HANDLE)
	coordScreen := COORD{0, 0}
	var dwConSize DWORD
	var csbi *CONSOLE_SCREEN_BUFFER_INFO
	csbi = GetConsoleScreenBufferInfo(hConsole)
	dwConSize = DWORD(csbi.DwSize.X * csbi.DwSize.Y)
	FillConsoleOutputCharacter(hConsole, TCHAR(' '), dwConSize, coordScreen)
	csbi = GetConsoleScreenBufferInfo(hConsole)
	FillConsoleOutputAttribute(hConsole, csbi.WAttributes, dwConSize, coordScreen)
	SetConsoleCursorPosition(hConsole, coordScreen)
}
func Gotoxy(x, y int16) {
	loc := COORD{x, y}
	hConsole := GetStdHandle(STD_OUTPUT_HANDLE)
	SetConsoleCursorPosition(hConsole, loc)
}
func Textbackground(color int) {
	hOut := GetStdHandle(STD_OUTPUT_HANDLE)
	SetConsoleTextAttribute(hOut, WORD(color))
}
func ResetColor() {
	Textbackground(int(OrigSrWindowColor))
}
