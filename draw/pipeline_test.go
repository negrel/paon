package draw

import (
	"bytes"
	"testing"

	"github.com/negrel/paon/geometry"
	"github.com/stretchr/testify/require"
)

func dumpBufferSurface(bf BufferSurface) string {
	var buf bytes.Buffer
	bf.Dump(&buf)
	return buf.String()
}

func TestPipeline(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		pipeline := NewPipeline(DrawerFunc(func(surface Surface) {
			surface.Set(geometry.NewVec2D(0, 0), Cell{
				Style:   CellStyle{},
				Content: '1',
			})
		}))

		surface := NewBufferSurface(geometry.NewSize(10, 2))

		pipeline.Draw(surface)

		require.Equal(t, ""+
			"1         \n"+
			"          \n",
			dumpBufferSurface(surface),
		)
	})

	t.Run("Double", func(t *testing.T) {
		pipeline := NewPipeline(DrawerFunc(func(surface Surface) {
			surface.Set(geometry.NewVec2D(0, 0), Cell{
				Style:   CellStyle{},
				Content: '1',
			})
		}))

		pipeline.Pipe(func(surface Surface, next Drawer) {
			// Will be overwritten by next drawer.
			surface.Set(geometry.NewVec2D(0, 0), Cell{
				Style:   CellStyle{},
				Content: '2',
			})
			next.Draw(surface)

			surface.Set(geometry.NewVec2D(1, 0), Cell{
				Style:   CellStyle{},
				Content: '2',
			})
		})

		surface := NewBufferSurface(geometry.NewSize(10, 2))

		pipeline.Draw(surface)

		require.Equal(t, ""+
			"12        \n"+
			"          \n",
			dumpBufferSurface(surface),
		)
	})

	t.Run("Triple", func(t *testing.T) {
		pipeline := NewPipeline(DrawerFunc(func(surface Surface) {
			surface.Set(geometry.NewVec2D(0, 0), Cell{
				Style:   CellStyle{},
				Content: '1',
			})
		}))

		pipeline.Pipe(func(surface Surface, next Drawer) {
			// Will be overwritten by next drawer.
			surface.Set(geometry.NewVec2D(0, 0), Cell{
				Style:   CellStyle{},
				Content: '2',
			})
			next.Draw(surface)

			surface.Set(geometry.NewVec2D(1, 0), Cell{
				Style:   CellStyle{},
				Content: '2',
			})
		})

		pipeline.Pipe(func(surface Surface, next Drawer) {
			// Will be overwritten by next drawer.
			surface.Set(geometry.NewVec2D(0, 0), Cell{
				Style:   CellStyle{},
				Content: '3',
			})
			surface.Set(geometry.NewVec2D(1, 0), Cell{
				Style:   CellStyle{},
				Content: '3',
			})

			next.Draw(surface)

			surface.Set(geometry.NewVec2D(2, 0), Cell{
				Style:   CellStyle{},
				Content: '3',
			})
		})

		surface := NewBufferSurface(geometry.NewSize(10, 2))

		pipeline.Draw(surface)

		require.Equal(t, ""+
			"123       \n"+
			"          \n",
			dumpBufferSurface(surface),
		)
	})
}
