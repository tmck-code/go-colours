package Colours

import (
	"fmt"
)

// List of ANSI colour codes
// source: https://bluesock.org/~willkg/dev/ansi.html
type Colour struct {
	Name string
	Fmt  string
}

var (
	Red         = Colour{Fmt: "[0;31m", Name: "Red"}
	Black       = Colour{Fmt: "[0;30m", Name: "Black"}
	Green       = Colour{Fmt: "[0;32m", Name: "Green"}
	Brown       = Colour{Fmt: "[0;33m", Name: "Brown"}
	Blue        = Colour{Fmt: "[0;34m", Name: "Blue"}
	Purple      = Colour{Fmt: "[0;35m", Name: "Purple"}
	Cyan        = Colour{Fmt: "[0;36m", Name: "Cyan"}
	Gray        = Colour{Fmt: "[0;37m", Name: "Gray"}
	DarkGray    = Colour{Fmt: "[1;30m", Name: "DarkGray"}
	LightRed    = Colour{Fmt: "[1;31m", Name: "LightRed"}
	LightGreen  = Colour{Fmt: "[1;32m", Name: "LightGreen"}
	Yellow      = Colour{Fmt: "[1;33m", Name: "Yellow"}
	LightBlue   = Colour{Fmt: "[1;34m", Name: "LightBlue"}
	LightPurple = Colour{Fmt: "[1;35m", Name: "LightPurple"}
	LightCyan   = Colour{Fmt: "[1;36m", Name: "LightCyan"}
	White       = Colour{Fmt: "[1;37m", Name: "White"}

	All = []Colour{
		Red,
		Black,
		Green,
		Brown,
		Blue,
		Purple,
		Cyan,
		Gray,
		DarkGray,
		LightRed,
		LightGreen,
		Yellow,
		LightBlue,
		LightPurple,
		LightCyan,
		White,
	}
)

func CPrintln(c Colour, s string) {
	fmt.Printf("\033%s%v\033[0m\n", c.Fmt, s)
}
func CPrint(c Colour, s string) {
	fmt.Printf("\033%s%v\033[0m", c.Fmt, s)
}
