package main

import (
	"fmt"
	"github.com/tmck-code/go-colours/colours"
)

func CPrintln(c Colour.Colour, s string) {
	fmt.Printf("\033%s%s\033[0m\n", c.Fmt, s)
}

func main() {
	CPrintln(Colour.Red, "hello world")
	// CPrintln(Black, "hello world")
	// CPrintln(Green, "hello world")
	// CPrintln(Brown, "hello world")
	// CPrintln(Blue, "hello world")
	// CPrintln(Purple, "hello world")
	// CPrintln(Cyan, "hello world")
	// CPrintln(Gray, "hello world")
	// CPrintln(DarkGray, "hello world")
	// CPrintln(LightRed, "hello world")
	// CPrintln(LightGreen, "hello world")
	// CPrintln(Yellow, "hello world")
	// CPrintln(LightBlue, "hello world")
	// CPrintln(LightPurple, "hello world")
	// CPrintln(LightCyan, "hello world")
	// CPrintln(White, "hello world")
}

// LightBlue = "\033[1;34m%s\033[0m"
