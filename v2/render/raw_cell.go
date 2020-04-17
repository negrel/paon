package render

import "github.com/gdamore/tcell"

// RawCell are ready to use cell for the painter.
type RawCell struct {
	X, Y  int
	Mainc rune
	Combc []rune
	Style tcell.Style
}
