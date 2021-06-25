package tcell

import (
	"github.com/gdamore/tcell/v2"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/styles/value"
)

var defaultFg, defaultBg, defaultAttr = tcell.StyleDefault.Decompose()

func fromTcell(mainc rune, combc []rune, style tcell.Style, width int) draw.Cell {
	fg, bg, attrs := style.Decompose()

	cell := draw.Cell{
		Style: draw.CellStyle{
			Bold:      (attrs & tcell.AttrBold) != 0,
			Blink:     (attrs & tcell.AttrBlink) != 0,
			Reverse:   (attrs & tcell.AttrReverse) != 0,
			Underline: (attrs & tcell.AttrUnderline) != 0,
			Dim:       (attrs & tcell.AttrDim) != 0,
			Italic:    (attrs & tcell.AttrItalic) != 0,
		},
		Content: mainc,
	}

	if fg != defaultFg {
		cell.Style.Foreground = value.ColorFromHex(fg.Hex())
	}
	if bg != defaultBg {
		cell.Style.Background = value.ColorFromHex(bg.Hex())
	}

	return cell
}

func toTcell(cell draw.Cell) (mainc rune, combc []rune, style tcell.Style) {
	style = tcell.StyleDefault.
		Bold(cell.Style.Bold).
		Blink(cell.Style.Blink).
		Reverse(cell.Style.Reverse).
		Underline(cell.Style.Underline).
		Dim(cell.Style.Dim).
		Italic(cell.Style.Italic)

	if cell.Style.Foreground.IsSet() {
		style = style.Foreground(
			tcell.NewHexColor(cell.Style.Foreground.Int32()),
		)
	}

	if cell.Style.Background.IsSet() {
		style = style.Background(
			tcell.NewHexColor(cell.Style.Background.Int32()),
		)
	}

	return cell.Content, []rune{}, style
}
