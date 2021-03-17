package tcell

import (
	"github.com/gdamore/tcell/v2"
	"github.com/negrel/paon/pkg/pdk/render"
)

func tcellStyle(cs render.CellStyle) tcell.Style {
	var foreground, background tcell.Color

	if cs.Foreground.IsSet() {
		foreground = tcell.NewHexColor(cs.Foreground.Int32())
	}
	if cs.Background.IsSet() {
		background = tcell.NewHexColor(cs.Background.Int32())
	}

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
