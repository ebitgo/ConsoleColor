# Go console text color

*Description: This library for windows and Linux console output font color control.*

## Installation


```shell
go get github.com/Ledgercn/ConsoleColor
```

In the Linux environment, call the fatih/color library.


```shell
go get github.com/fatih/color
```
## Usage


```go
func Example() {
	ConsoleColor.Print(ConsoleColor.C_RED, "This is a test!\r\n")
	ConsoleColor.Printf(ConsoleColor.C_BLUE, "%s!\r\n", "This is a test!")
	ConsoleColor.Println(ConsoleColor.C_WHITE, "This is a test!")

	ConsoleColor.Println(ConsoleColor.C_BLACK, "This is black")
	ConsoleColor.Println(ConsoleColor.C_RED, "This is red")
	ConsoleColor.Println(ConsoleColor.C_GREEN, "This is green")
	ConsoleColor.Println(ConsoleColor.C_YELLOW, "This is yellow")
	ConsoleColor.Println(ConsoleColor.C_BLUE, "This is blue")
	ConsoleColor.Println(ConsoleColor.C_MAGENTA, "This is megenta")
	ConsoleColor.Println(ConsoleColor.C_CYAN, "This is cyan")
	ConsoleColor.Println(ConsoleColor.C_WHITE, "This is white")
}
```