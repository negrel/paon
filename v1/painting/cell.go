package painting

import (
	"github.com/negrel/ginger/v1/color"
)

// CellDefault is the default cell used for
// Matrix constructor.
var CellDefault = Cell{
	Style: color.StyleDefault,
	Char:  0,
}

// CellOverflow is used when a component is too large
// to fit the constraint.
var CellOverflow = Cell{
	Style: color.Style{
		Foreground: 0xFFFFFF,
		Background: 0xFF0000,
	},
	Char: '!',
}

// Cell is an element in the terminal screen matrix.
type Cell struct {
	color.Style
	Char rune
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Compute return the raw cell for the painter.
func (c *Cell) Compute(p Position) RawCell {

	return RawCell{
		Position: p,
		Mainc:    c.Char,
		Style:    c.Style.Compute(),
	}
}

func (c *Cell) isEqual(other *Cell) bool {
	if c.Char != other.Char ||
		c.Foreground != other.Foreground ||
		c.Background != other.Background {

		return false
	}

	return true
}
