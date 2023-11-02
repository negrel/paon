package tcell

import (
	"math/rand"
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/negrel/paon/draw"
)

func randomCellStyle() draw.CellStyle {
	return draw.CellStyle{
		Foreground:    randColor(),
		Background:    randColor(),
		Bold:          randPropBool(),
		Blink:         randPropBool(),
		Dim:           randPropBool(),
		Reverse:       randPropBool(),
		Underline:     randPropBool(),
		Italic:        randPropBool(),
		StrikeThrough: randPropBool(),
	}
}

func randomCell() draw.Cell {
	return draw.Cell{
		Style:   randomCellStyle(),
		Content: randRune(),
	}
}

func randomTcellStyle() tcell.Style {
	return tcell.StyleDefault.
		Foreground(tcell.NewHexColor(rand.Int31())).
		Background(tcell.NewHexColor(rand.Int31())).
		Blink(randBool()).
		Bold(randBool()).
		Dim(randBool()).
		Italic(randBool()).
		Reverse(randBool()).
		Underline(randBool()).
		StrikeThrough(randBool())
}

func randomTcellCell() (rune, []rune, tcell.Style) {
	return randRune(), []rune{}, randomTcellStyle()
}

func BenchmarkToTcell(b *testing.B) {
	cell := randomCell()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mainc, combc, style := toTcell(cell)
		_ = mainc
		_ = combc
		_ = style
	}
}

func BenchmarkFromTcell(b *testing.B) {
	mainc, combc, style := randomTcellCell()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cell := fromTcell(mainc, combc, style, 1)
		_ = cell
	}
}
