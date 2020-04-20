package render

import "github.com/negrel/ginger/v2/style"

// Cell represents a cell of the terminal screen.
type Cell struct {
	Char  rune
	Theme style.Theme
}

// CellDefault is the default cell used for
// Matrix constructor.
var CellDefault = Cell{
	Char:  ' ',
	Theme: style.DefaultTheme,
}

// CellOverflow is used when a component is too large
// to fit the constraint.
var CellOverflow = Cell{
	Char:  '!',
	Theme: style.DefaultTheme.Background(0xFF0000).Foreground(0xFFFFFF).Blink(true),
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Compute return a raw cell ready to be drawn on the
// screen.
func (c *Cell) Compute(p Position) *RawCell {
	return &RawCell{
		X:     p.X,
		Y:     p.Y,
		Mainc: c.Char,
		Style: c.Theme.Compute(),
	}
}

func (c *Cell) isEqual(other *Cell) bool {
	if c.Char != other.Char ||
		c.Theme != other.Theme {

		return false
	}

	return true
}
