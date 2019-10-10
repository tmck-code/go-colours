package main

import (
	"fmt"
	"github.com/tmck-code/go-colours/colours"
)

func CPrintln(c Colour.Colour, s string) {
	fmt.Printf("%-15s \033%s%s\033[0m\n", c.Name, c.Fmt, s)
}

func CPrintVariantln(c Colour.Colour, variant string, s string) {
	// fmt.Println(c.Name, variant, c.Fmt, s)
	fmt.Printf("%-15s %s \033[%s;%sm %s \033[0m\n", c.Name, variant, variant, c.Fmt, s)
}

func main() {
	for _, c := range Colour.All {
		// CPrintln(c, "hello world")
		CPrintVariantln(c, "0", "hello world")
		CPrintVariantln(c, "1", "hello world")
		CPrintVariantln(c, "2", "hello world")
		CPrintVariantln(c, "3", "hello world")
		CPrintVariantln(c, "4", "hello world")
		CPrintVariantln(c, "7", "hello world")
		CPrintVariantln(c, "9", "hello world")
	}
}
