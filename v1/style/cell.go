package style

import (
	"github.com/gdamore/tcell"
)

// Colors define the foreground & background color.
type Colors = tcell.Style

// Cell represent a terminal screen cell.
type Cell struct {
	Char   rune
	Colors Colors
}

// Row is a row of terminal screen cell.
type Row []*Cell

// Col is a column of terminal screen cell.
type Col = Row

// NewRow return a row of the given size.
func NewRow(len int, c *Cell) Row {
	r := make(Row, len)

	for i := 0; i < len; i++ {
		r[i] = c
	}

	return r
}
