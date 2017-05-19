package ConsoleColor

import (
	"runtime"
)

const (
	OS_UNKNOWN = iota
	OS_WINDOWS
	OS_LINUX
	OS_MACOSX
)

func GetOS() int {
	switch runtime.GOOS {
	case "windows":
		return OS_WINDOWS
	case "linux":
		return OS_LINUX
	case "darwin":
		return OS_MACOSX
	default:
		return OS_UNKNOWN
	}
}
