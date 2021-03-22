package tcell

import (
	"github.com/gdamore/tcell/v2"
	"github.com/negrel/paon/pkg/pdk/render"
)

func tcellStyle(cs render.CellStyle) tcell.Style {
	foreground := tcell.NewHexColor(cs.Foreground.Int32())
	background := tcell.NewHexColor(cs.Background.Int32())

	return tcell.StyleDefault.
		Foreground(foreground).
		Background(background).
		Bold(cs.Bold).
		Blink(cs.Blink).
		Reverse(cs.Reverse).
		Underline(cs.Underline).
		Dim(cs.Dim).
		Italic(cs.Italic).
		StrikeThrough(cs.StrikeThrough)
}
