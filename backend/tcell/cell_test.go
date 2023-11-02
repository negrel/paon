package tcell

import (
	"math/rand"
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/negrel/paon/colors"
	"github.com/negrel/paon/draw"
	"github.com/stretchr/testify/require"
)

func foreground(style tcell.Style) tcell.Color {
	fg, _, _ := style.Decompose()
	return fg
}

func background(style tcell.Style) tcell.Color {
	_, bg, _ := style.Decompose()
	return bg
}

func blink(style tcell.Style) bool {
	_, _, attrs := style.Decompose()
	return attrs&tcell.AttrBlink != 0
}

func bold(style tcell.Style) bool {
	_, _, attrs := style.Decompose()
	return attrs&tcell.AttrBold != 0
}

func dim(style tcell.Style) bool {
	_, _, attrs := style.Decompose()
	return attrs&tcell.AttrDim != 0
}

func italic(style tcell.Style) bool {
	_, _, attrs := style.Decompose()
	return attrs&tcell.AttrItalic != 0
}

func reverse(style tcell.Style) bool {
	_, _, attrs := style.Decompose()
	return attrs&tcell.AttrReverse != 0
}

func underline(style tcell.Style) bool {
	_, _, attrs := style.Decompose()
	return attrs&tcell.AttrUnderline != 0
}

func strikethrough(style tcell.Style) bool {
	_, _, attrs := style.Decompose()
	return attrs&tcell.AttrStrikeThrough != 0
}

func randColor() colors.Color {
	return colors.ColorFromHex(rand.Uint32())
}

func randBool() bool {
	return bool(randPropBool())
}

func randPropBool() bool {
	return rand.Int()%2 == 0
}

func randRune() rune {
	return rune(rand.Int31())
}

func TestToCellStyle(t *testing.T) {
	// Default
	style := toTcellStyle(draw.CellStyle{})
	require.Equal(t, foreground(tcell.StyleDefault), foreground(style))
	require.Equal(t, background(tcell.StyleDefault), background(style))
	require.Equal(t, blink(tcell.StyleDefault), blink(style))
	require.Equal(t, bold(tcell.StyleDefault), bold(style))
	require.Equal(t, dim(tcell.StyleDefault), dim(style))
	require.Equal(t, italic(tcell.StyleDefault), italic(style))
	require.Equal(t, reverse(tcell.StyleDefault), reverse(style))
	require.Equal(t, underline(tcell.StyleDefault), underline(style))
	require.Equal(t, strikethrough(tcell.StyleDefault), strikethrough(style))

	// Colors
	cellstyle := draw.CellStyle{
		Background: randColor(),
		Foreground: randColor(),
	}

	style = toTcellStyle(cellstyle)
	require.Equal(t, cellstyle.Foreground.Hex(), uint32(foreground(style).Hex()))
	require.Equal(t, cellstyle.Background.Hex(), uint32(background(style).Hex()))

	// Attributes
	cellstyle = draw.CellStyle{
		Blink:         randPropBool(),
		Bold:          randPropBool(),
		Dim:           randPropBool(),
		Italic:        randPropBool(),
		Reverse:       randPropBool(),
		Underline:     randPropBool(),
		StrikeThrough: randPropBool(),
	}

	style = toTcellStyle(cellstyle)
	require.EqualValues(t, cellstyle.Blink, blink(style))
	require.EqualValues(t, cellstyle.Bold, bold(style))
	require.EqualValues(t, cellstyle.Dim, dim(style))
	require.EqualValues(t, cellstyle.Italic, italic(style))
	require.EqualValues(t, cellstyle.Reverse, reverse(style))
	require.EqualValues(t, cellstyle.Underline, underline(style))
	require.EqualValues(t, cellstyle.StrikeThrough, strikethrough(style))
}

func TestToCell_Content(t *testing.T) {
	cell := draw.Cell{
		Content: randRune(),
	}

	mainc, _, _ := toTcell(cell)
	require.Equal(t, cell.Content, mainc)
}

func TestFromTcellStyle(t *testing.T) {
	// Default
	cellstyle := fromTcellStyle(tcell.StyleDefault)

	require.EqualValues(t, cellstyle.Foreground, colors.ColorUnset())
	require.EqualValues(t, cellstyle.Background, colors.ColorUnset())
	require.EqualValues(t, cellstyle.Blink, blink(tcell.StyleDefault))
	require.EqualValues(t, cellstyle.Bold, bold(tcell.StyleDefault))
	require.EqualValues(t, cellstyle.Dim, dim(tcell.StyleDefault))
	require.EqualValues(t, cellstyle.Italic, italic(tcell.StyleDefault))
	require.EqualValues(t, cellstyle.Reverse, reverse(tcell.StyleDefault))
	require.EqualValues(t, cellstyle.Underline, underline(tcell.StyleDefault))
	require.EqualValues(t, cellstyle.StrikeThrough, strikethrough(tcell.StyleDefault))

	// Colors
	style := tcell.StyleDefault.
		Foreground(tcell.NewHexColor(rand.Int31())).
		Background(tcell.NewHexColor(rand.Int31()))

	cellstyle = fromTcellStyle(style)
	require.Equal(t, cellstyle.Foreground, fromTcellColor(foreground(style)))
	require.Equal(t, cellstyle.Background, fromTcellColor(background(style)))
}

func TestFromCellContent(t *testing.T) {
	mainc := randRune()
	cell := fromTcell(mainc, []rune{}, tcell.StyleDefault, 1)
	require.Equal(t, cell.Content, mainc)
}
