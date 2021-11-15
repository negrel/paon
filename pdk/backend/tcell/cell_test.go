package tcell

import (
	"math/rand"
	"testing"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/styles/property"
	"github.com/stretchr/testify/assert"
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

func randColor() property.Color {
	return property.ColorFromHex(rand.Uint32())
}

func randBool() bool {
	return bool(randPropBool())
}

func randPropBool() property.Bool {
	return rand.Int()%2 == 0
}

func randRune() rune {
	return rune(rand.Int31())
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestToCellStyle(t *testing.T) {
	// Default
	style := toTcellStyle(draw.CellStyle{})
	assert.Equal(t, foreground(tcell.StyleDefault), foreground(style))
	assert.Equal(t, background(tcell.StyleDefault), background(style))
	assert.Equal(t, blink(tcell.StyleDefault), blink(style))
	assert.Equal(t, bold(tcell.StyleDefault), bold(style))
	assert.Equal(t, dim(tcell.StyleDefault), dim(style))
	assert.Equal(t, italic(tcell.StyleDefault), italic(style))
	assert.Equal(t, reverse(tcell.StyleDefault), reverse(style))
	assert.Equal(t, underline(tcell.StyleDefault), underline(style))
	assert.Equal(t, strikethrough(tcell.StyleDefault), strikethrough(style))

	// Colors
	cellstyle := draw.CellStyle{
		Background: randColor(),
		Foreground: randColor(),
	}

	style = toTcellStyle(cellstyle)
	assert.Equal(t, cellstyle.Foreground.Hex(), uint32(foreground(style).Hex()))
	assert.Equal(t, cellstyle.Background.Hex(), uint32(background(style).Hex()))

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
	assert.Equal(t, cellstyle.Blink, blink(style))
	assert.Equal(t, cellstyle.Bold, bold(style))
	assert.Equal(t, cellstyle.Dim, dim(style))
	assert.Equal(t, cellstyle.Italic, italic(style))
	assert.Equal(t, cellstyle.Reverse, reverse(style))
	assert.Equal(t, cellstyle.Underline, underline(style))
	assert.Equal(t, cellstyle.StrikeThrough, strikethrough(style))
}

func TestToCell_Content(t *testing.T) {
	cell := draw.Cell{
		Content: randRune(),
	}

	mainc, _, _ := toTcell(cell)
	assert.Equal(t, cell.Content, mainc)
}

func TestFromTcellStyle(t *testing.T) {
	// Default
	cellstyle := fromTcellStyle(tcell.StyleDefault)

	assert.Equal(t, cellstyle.Foreground, property.ColorUnset())
	assert.Equal(t, cellstyle.Background, property.ColorUnset())
	assert.Equal(t, cellstyle.Blink, blink(tcell.StyleDefault))
	assert.Equal(t, cellstyle.Bold, bold(tcell.StyleDefault))
	assert.Equal(t, cellstyle.Dim, dim(tcell.StyleDefault))
	assert.Equal(t, cellstyle.Italic, italic(tcell.StyleDefault))
	assert.Equal(t, cellstyle.Reverse, reverse(tcell.StyleDefault))
	assert.Equal(t, cellstyle.Underline, underline(tcell.StyleDefault))
	assert.Equal(t, cellstyle.StrikeThrough, strikethrough(tcell.StyleDefault))

	// Colors
	style := tcell.StyleDefault.
		Foreground(tcell.NewHexColor(rand.Int31())).
		Background(tcell.NewHexColor(rand.Int31()))

	cellstyle = fromTcellStyle(style)
	assert.Equal(t, cellstyle.Foreground, fromTcellColor(foreground(style)))
	assert.Equal(t, cellstyle.Background, fromTcellColor(background(style)))
}

func TestFromCell_Content(t *testing.T) {
	mainc := randRune()
	cell := fromTcell(mainc, []rune{}, tcell.StyleDefault, 1)
	assert.Equal(t, cell.Content, mainc)
}
