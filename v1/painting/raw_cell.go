package painting

import "github.com/gdamore/tcell"

// RawCell are ready to use cell for the painter.
type RawCell struct {
	Position
	Mainc rune
	Combc []rune
	Style tcell.Style
}
