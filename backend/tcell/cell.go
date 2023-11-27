package tcell

import (
	"github.com/gdamore/tcell/v2"
	"github.com/negrel/paon/colors"
	"github.com/negrel/paon/draw"
)

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
		cellstyle.Foreground = fromTcellColor(fg)
	}
	if bg != defaultBg {
		cellstyle.Background = fromTcellColor(bg)
	}

	return cellstyle
}

func toTcell(cell draw.Cell) (rune, []rune, tcell.Style) {
	return cell.Content, []rune{}, toTcellStyle(cell.Style)
}

func toTcellStyle(cellstyle draw.CellStyle) tcell.Style {
	style := tcell.StyleDefault.
		Bold(bool(cellstyle.Bold)).
		Blink(bool(cellstyle.Blink)).
		Reverse(bool(cellstyle.Reverse)).
		Underline(bool(cellstyle.Underline)).
		Dim(bool(cellstyle.Dim)).
		Italic(bool(cellstyle.Italic)).
		StrikeThrough(bool(cellstyle.StrikeThrough))

	if cellstyle.Foreground.A() != 0 {
		style = style.Foreground(
			toTcellColor(cellstyle.Foreground),
		)
	}

	if cellstyle.Background.A() != 0 {
		style = style.Background(
			toTcellColor(cellstyle.Background),
		)
	}

	return style
}

func toTcellColor(color colors.Color) tcell.Color {
	return tcell.NewHexColor(int32(color.Hex()) & 0xFFFFFF)
}

func fromTcellColor(color tcell.Color) colors.Color {
	return colors.ColorFromHex(uint32(color.Hex()) & 0xFFFFFF)
}
