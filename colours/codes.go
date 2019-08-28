package Colour

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
	All         = []Colour{
		Red, Green,
	}
)
