package render

import "github.com/gdamore/tcell"

// RawCell are ready to use cell.
type RawCell struct {
	X, Y  int
	Mainc rune
	Style tcell.Style
}
