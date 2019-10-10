package Colour

// List of ANSI colour codes
// source: https://bluesock.org/~willkg/dev/ansi.html
type Colour struct {
	Name string
	Fmt  string
}

var (
	Black  = Colour{Fmt: "30", Name: "Black"}
	Red    = Colour{Fmt: "31", Name: "Red"}
	Green  = Colour{Fmt: "32", Name: "Green"}
	Brown  = Colour{Fmt: "33", Name: "Brown"}
	Blue   = Colour{Fmt: "34", Name: "Blue"}
	Purple = Colour{Fmt: "35", Name: "Purple"}
	Cyan   = Colour{Fmt: "36", Name: "Cyan"}
	Gray   = Colour{Fmt: "37", Name: "Gray"}
	All    = []Colour{
		Red,
		Black,
		Green,
		Brown,
		Blue,
		Purple,
		Cyan,
		Gray,
	}
)
