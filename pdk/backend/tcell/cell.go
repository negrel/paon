package tcell

import (
	"github.com/gdamore/tcell/v2"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/styles/value"
)

var _, _, defaultAttr = tcell.StyleDefault.Decompose()
var defaultFg = tcell.ColorDefault
var defaultBg = tcell.ColorDefault

func fromTcell(mainc rune, combc []rune, style tcell.Style, width int) draw.Cell {
	cell := draw.Cell{
		Style:   fromTcellStyle(style),
		Content: mainc,
	}

	return cell
}

func fromTcellStyle(style tcell.Style) draw.CellStyle {
	fg, bg, attrs := style.Decompose()
	cellstyle := draw.CellStyle{
		Bold:      (attrs & tcell.AttrBold) != 0,
		Blink:     (attrs & tcell.AttrBlink) != 0,
		Reverse:   (attrs & tcell.AttrReverse) != 0,
		Underline: (attrs & tcell.AttrUnderline) != 0,
		Dim:       (attrs & tcell.AttrDim) != 0,
		Italic:    (attrs & tcell.AttrItalic) != 0,
	}

	if fg != defaultFg {
		cellstyle.Foreground = value.ColorFromHex(fg.Hex())
	}
	if bg != defaultBg {
		cellstyle.Background = value.ColorFromHex(bg.Hex())
	}

	return cellstyle
}

func toTcell(cell draw.Cell) (rune, []rune, tcell.Style) {
	return cell.Content, []rune{}, toTcellStyle(cell.Style)
}

func toTcellStyle(cellstyle draw.CellStyle) tcell.Style {
	style := tcell.StyleDefault.
		Bold(cellstyle.Bold).
		Blink(cellstyle.Blink).
		Reverse(cellstyle.Reverse).
		Underline(cellstyle.Underline).
		Dim(cellstyle.Dim).
		Italic(cellstyle.Italic).
		StrikeThrough(cellstyle.StrikeThrough)

	if cellstyle.Foreground.IsSet() {
		style = style.Foreground(
			tcell.NewHexColor(cellstyle.Foreground.Hex()),
		)
	}

	if cellstyle.Background.IsSet() {
		style = style.Background(
			tcell.NewHexColor(cellstyle.Background.Hex()),
		)
	}

	return style
}
