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
		bf := NewBufferSurface(geometry.NewSize(10, 2))
		bf.Set(geometry.NewVec2D(0, 0), Cell{
			Style:   CellStyle{},
			Content: 'a',
		})
		bf.Set(geometry.NewVec2D(9, 0), Cell{
			Style:   CellStyle{},
			Content: 'b',
		})
		bf.Set(geometry.NewVec2D(0, 1), Cell{
			Style:   CellStyle{},
			Content: 'c',
		})
		bf.Set(geometry.NewVec2D(9, 1), Cell{
			Style:   CellStyle{},
			Content: 'd',
		})

		var buf bytes.Buffer
		bf.Dump(&buf)

		require.Equal(t, ""+
			"a        b\n"+
			"c        d\n",
			buf.String(),
		)
	})

	t.Run("OutOfBoundDrawing", func(t *testing.T) {
		bf := NewBufferSurface(geometry.NewSize(10, 1))

		bf.Set(geometry.NewVec2D(0, 1), Cell{
			Style:   CellStyle{},
			Content: '!',
		})

		var buf bytes.Buffer
		bf.Dump(&buf)

		require.Regexp(t, regexp.MustCompile(`\s{10}\n`), buf.String())
	})

	t.Run("Overwrite", func(t *testing.T) {
		bf := NewBufferSurface(geometry.NewSize(10, 1))

		bf.Set(geometry.NewVec2D(0, 0), Cell{
			Style:   CellStyle{},
			Content: '!',
		})
		bf.Set(geometry.NewVec2D(0, 0), Cell{
			Style:   CellStyle{},
			Content: '.',
		})

		var buf bytes.Buffer
		bf.Dump(&buf)

		require.Regexp(t, regexp.MustCompile(`\.\s{9}\n`), buf.String())
	})
}
