package render

import (
	"github.com/gdamore/tcell/v2"
	"github.com/negrel/paon/pkg/pdk/styles/value"
)

type CellStyle struct {
	Foreground    value.Color
	Background    value.Color
	Bold          bool
	Blink         bool
	Reverse       bool
	Underline     bool
	Dim           bool
	Italic        bool
	StrikeThrough bool
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

// Cell define a terminal Screen cell.
type Cell struct {
	Style   CellStyle
	Content rune
}

func makeCellFromTcell(content rune, style tcell.Style) Cell {
	fg, bg, attr := style.Decompose()

	cstyle := CellStyle{
		Foreground:    value.ColorFromHex(int32(fg)),
		Background:    value.ColorFromHex(int32(bg)),
		Bold:          (tcell.AttrBold & attr) != 0,
		Blink:         (tcell.AttrBlink & attr) != 0,
		Reverse:       (tcell.AttrReverse & attr) != 0,
		Underline:     (tcell.AttrUnderline & attr) != 0,
		Dim:           (tcell.AttrDim & attr) != 0,
		Italic:        (tcell.AttrItalic & attr) != 0,
		StrikeThrough: (tcell.AttrStrikeThrough & attr) != 0,
	}

	return Cell{
		Style:   cstyle,
		Content: content,
	}
}
