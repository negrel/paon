package draw

import (
	"github.com/gdamore/tcell/v2"
)

type CellStyle tcell.Style

// Cell define a terminal Screen cell.
type Cell struct {
	Style   CellStyle
	Content rune
}

func makeCellFromTcell(content rune, style tcell.Style) Cell {
	return Cell{
		Style:   CellStyle(style),
		Content: content,
	}
}
