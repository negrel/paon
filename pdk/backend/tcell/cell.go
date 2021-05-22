package tcell

import (
	"github.com/gdamore/tcell/v2"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/styles/value"
)

func fromTcell(mainc rune, combc []rune, style tcell.Style, width int) draw.Cell {
	fg, bg, attrs := style.Decompose()

	return draw.Cell{
		Style: draw.CellStyle{
			Foreground: value.ColorFromHex(fg.Hex()),
			Background: value.ColorFromHex(bg.Hex()),
			Bold:       (attrs & tcell.AttrBold) != 0,
			Blink:      (attrs & tcell.AttrBlink) != 0,
			Reverse:    (attrs & tcell.AttrReverse) != 0,
			Underline:  (attrs & tcell.AttrUnderline) != 0,
			Dim:        (attrs & tcell.AttrDim) != 0,
			Italic:     (attrs & tcell.AttrItalic) != 0,
		},
		Content: mainc,
	}
}

func toTcell(cell draw.Cell) (mainc rune, combc []rune, style tcell.Style) {
	return cell.Content, []rune{},
		tcell.StyleDefault.
			Foreground(
				tcell.NewHexColor(cell.Style.Foreground.Int32()),
			).
			Background(
				tcell.NewHexColor(cell.Style.Background.Int32()),
			).
			Bold(cell.Style.Bold).
			Blink(cell.Style.Blink).
			Reverse(cell.Style.Reverse).
			Underline(cell.Style.Underline).
			Dim(cell.Style.Dim).
			Italic(cell.Style.Italic)
}
