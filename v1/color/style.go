package color

import "github.com/gdamore/tcell"

// Style define the color style of a cell.
type Style struct {
	Foreground Color
	Background Color
}

// StyleDefault is used to leave the foreground
// and the background color unchanged from
// whatever system or teminal default may exist.
var StyleDefault = Style{
	Foreground: Default,
	Background: Default,
}

// Compute method return the style value ready
// for the painting.
func (s *Style) Compute() tcell.Style {
	st := tcell.StyleDefault

	if s.Foreground != Default {
		st = st.Background(tcell.NewHexColor(s.Background))
	}

	if s.Background != Default {
		st = st.Foreground(tcell.NewHexColor(s.Foreground))
	}

	return st
}
