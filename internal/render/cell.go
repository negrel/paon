package render

import "github.com/gdamore/tcell"

// Cell define a terminal screen cell.
type Cell struct {
	style   *tcell.Style
	content rune
}
