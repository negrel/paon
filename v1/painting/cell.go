package painting

import (
	"github.com/gdamore/tcell"
)

// CellDefault is the default cell used for
// Matrix constructor.
var CellDefault = Cell{
	Foreground: tcell.ColorDefault,
	Background: tcell.ColorDefault,
}

// Cell is an element in the terminal screen matrix.
type Cell struct {
	Char       rune
	Foreground tcell.Color
	Background tcell.Color
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
		Style:    tcell.StyleDefault.Background(c.Background).Foreground(c.Foreground),
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
