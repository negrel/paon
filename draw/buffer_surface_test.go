package draw

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/negrel/paon/geometry"
	"github.com/stretchr/testify/require"
)

func TestBufferSurface(t *testing.T) {
	t.Run("CornerDrawing", func(t *testing.T) {
		bf := NewBufferSurface(geometry.Size{Width: 10, Height: 2})
		bf.Set(geometry.Vec2D{X: 0, Y: 0}, Cell{
			Style:   CellStyle{},
			Content: 'a',
		})
		bf.set(geometry.Vec2D{X: 9, Y: 0}, Cell{Style: CellStyle{}, Content: 'b'})
		bf.Set(geometry.Vec2D{X: 0, Y: 1}, Cell{
			Style:   CellStyle{},
			Content: 'c',
		})
		bf.set(geometry.Vec2D{X: 9, Y: 1}, Cell{Style: CellStyle{}, Content: 'd'})

		var buf bytes.Buffer
		err := bf.Dump(&buf)
		require.NoError(t, err)

		require.Equal(t, ""+
			"a        b\n"+
			"c        d\n",
			buf.String(),
		)
	})

	t.Run("OutOfBoundDrawing", func(t *testing.T) {
		bf := NewBufferSurface(geometry.Size{Width: 10, Height: 1})

		bf.Set(geometry.Vec2D{X: 0, Y: 1}, Cell{
			Style:   CellStyle{},
			Content: '!',
		})

		var buf bytes.Buffer
		err := bf.Dump(&buf)
		require.NoError(t, err)

		require.Regexp(t, regexp.MustCompile(`\s{10}\n`), buf.String())
	})

	t.Run("Overwrite", func(t *testing.T) {
		bf := NewBufferSurface(geometry.Size{Width: 10, Height: 1})

		bf.Set(geometry.Vec2D{X: 0, Y: 0}, Cell{
			Style:   CellStyle{},
			Content: '!',
		})
		bf.Set(geometry.Vec2D{X: 0, Y: 0}, Cell{
			Style:   CellStyle{},
			Content: '.',
		})

		var buf bytes.Buffer
		err := bf.Dump(&buf)
		require.NoError(t, err)

		require.Regexp(t, regexp.MustCompile(`\.\s{9}\n`), buf.String())
	})
}
