package paint

import "github.com/negrel/ginger/v1/style"

// RawCell is the raw information needed by the
// painter to paint a cell.
type RawCell struct {
	X, Y  int
	Mainc rune
	Conbc []rune
	Style style.Colors
}
