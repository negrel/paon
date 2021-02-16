package render

import (
	"github.com/gdamore/tcell/v2"
	"github.com/negrel/paon/pkg/pdk/styles/value"
)

type CellStyle struct {
	Foreground value.Color
	Background value.Color

	Bold, Blink, Reverse, Underline,
	Dim, Italic, StrikeThrough bool
}

func (cs CellStyle) toTcellStyle() tcell.Style {
	return tcell.StyleDefault.
		Foreground(tcell.Color(cs.Foreground.Int32())).
		Background(tcell.Color(cs.Background.Int32())).
		Bold(cs.Bold).
		Blink(cs.Blink).
		Reverse(cs.Reverse).
		Underline(cs.Underline).
		Dim(cs.Dim).
		Italic(cs.Italic).
		StrikeThrough(cs.StrikeThrough)
}

// Cell define a terminal Surface cell.
type Cell struct {
	Style   CellStyle
	Content rune
}
