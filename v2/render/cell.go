package render

import (
	"github.com/negrel/ginger/v2/style"
)

// CellDefault is the default cell used for
// Matrix constructor.
var CellDefault = Cell{
	Theme: &style.DefaultTheme,
	Char:  0,
}

var overflowTheme = style.DefaultTheme.Foreground(0xFFFFFF).Background(0xFF0000).Blink(true)

// CellOverflow is used when a component is too large
// to fit the constraint.
var CellOverflow = Cell{
	Theme: &overflowTheme,
	Char:  '!',
}

// Cell is an element in the terminal screen matrix.
type Cell struct {
	Theme *style.Theme
	Char  rune
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Compute return the raw cell for the painter.
func (c *Cell) Compute(p Position) *RawCell {
	return &RawCell{
		X:     p.X,
		Y:     p.Y,
		Mainc: c.Char,
		Style: *c.Theme,
	}
}

func (c *Cell) isEqual(other *Cell) bool {
	if c.Char != other.Char ||
		c.Theme != other.Theme {

		return false
	}

	return true
}
