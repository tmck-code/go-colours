package main

import "fmt"

// List of ANSI colour codes
// source: https://bluesock.org/~willkg/dev/ansi.html
type Colour struct {
	Name string
	Fmt  string
}

var (
	Red         = Colour{Name: "Red", Fmt: "[0;31m"}
	Black       = Colour{Name: "Black", Fmt: "[0;30m"}
	Green       = Colour{Name: "Green", Fmt: "[0;32m"}
	Brown       = Colour{Name: "Brown", Fmt: "[0;33m"}
	Blue        = Colour{Name: "Blue", Fmt: "[0;34m"}
	Purple      = Colour{Name: "Purple", Fmt: "[0;35m"}
	Cyan        = Colour{Name: "Cyan", Fmt: "[0;36m"}
	Gray        = Colour{Name: "Gray", Fmt: "[0;37m"}
	DarkGray    = Colour{Name: "DarkGray", Fmt: "[1;30m"}
	LightRed    = Colour{Name: "LightRed", Fmt: "[1;31m"}
	LightGreen  = Colour{Name: "LightGreen", Fmt: "[1;32m"}
	Yellow      = Colour{Name: "Yellow", Fmt: "[1;33m"}
	LightBlue   = Colour{Name: "LightBlue", Fmt: "[1;34m"}
	LightPurple = Colour{Name: "LightPurple", Fmt: "[1;35m"}
	LightCyan   = Colour{Name: "LightCyan", Fmt: "[1;36m"}
	White       = Colour{Name: "White", Fmt: "[1;37m"}
)

func CPrintln(c Colour, s string) {
	fmt.Printf("\033%s%s\033[0m\n", c.Fmt, s)
}

func main() {
	CPrintln(Red, "hello world")
	CPrintln(Black, "hello world")
	CPrintln(Green, "hello world")
	CPrintln(Brown, "hello world")
	CPrintln(Blue, "hello world")
	CPrintln(Purple, "hello world")
	CPrintln(Cyan, "hello world")
	CPrintln(Gray, "hello world")
	CPrintln(DarkGray, "hello world")
	CPrintln(LightRed, "hello world")
	CPrintln(LightGreen, "hello world")
	CPrintln(Yellow, "hello world")
	CPrintln(LightBlue, "hello world")
	CPrintln(LightPurple, "hello world")
	CPrintln(LightCyan, "hello world")
	CPrintln(White, "hello world")
}

// LightBlue = "\033[1;34m%s\033[0m"
