package main

import (
	"fmt"
	"github.com/tmck-code/go-colours/colours"
)

func CPrintln(c Colour.Colour, s string) {
	fmt.Printf("\033%s%s\033[0m\n", c.Fmt, s)
}

func main() {
	for _, c := range Colour.All {
		CPrintln(c, "hello world")
	}
}
