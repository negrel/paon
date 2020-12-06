package draw

import (
	"github.com/gdamore/tcell"
)

// Cell define a terminal screen cell.
type Cell struct {
	Style   tcell.Style
	Content rune
}
