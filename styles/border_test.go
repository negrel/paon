package styles

import (
	"testing"

	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/geometry"
	"github.com/stretchr/testify/require"
)

func TestBorderDrawer(t *testing.T) {
	t.Run("BorderHidden", func(t *testing.T) {
		t.Run("DefaultStyle", func(t *testing.T) {
			drawer := BorderDrawers[BorderHidden]
			style := BordersStyle{}
			surface := draw.NewBufferSurface(geometry.NewSize(10, 10))

			t.Run("Left", func(t *testing.T) {
				drawer.Left(style, surface)

				for x := 0; x < surface.Size().Width(); x++ {
					for y := 0; y < surface.Size().Height(); y++ {
						require.Equal(t, draw.Cell{}, surface.Get(geometry.NewVec2D(x, y)))
					}
				}
			})

			t.Run("Top", func(t *testing.T) {
				drawer.Top(style, surface)

				for x := 0; x < surface.Size().Width(); x++ {
					for y := 0; y < surface.Size().Height(); y++ {
						require.Equal(t, draw.Cell{}, surface.Get(geometry.NewVec2D(x, y)))
					}
				}
			})

			t.Run("Right", func(t *testing.T) {
				drawer.Right(style, surface)

				for x := 0; x < surface.Size().Width(); x++ {
					for y := 0; y < surface.Size().Height(); y++ {
						require.Equal(t, draw.Cell{}, surface.Get(geometry.NewVec2D(x, y)))
					}
				}
			})

			t.Run("Bottom", func(t *testing.T) {
				drawer.Bottom(style, surface)

				for x := 0; x < surface.Size().Width(); x++ {
					for y := 0; y < surface.Size().Height(); y++ {
						require.Equal(t, draw.Cell{}, surface.Get(geometry.NewVec2D(x, y)))
					}
				}
			})
		})

		t.Run("BorderHidden/DefaultCellStyle", func(t *testing.T) {
			drawer := BorderDrawers[BorderHidden]
			style := BordersStyle{
				Top: BorderSide{
					Size:      0,
					Style:     BorderHidden,
					CellStyle: draw.CellStyle{},
				},
				Bottom: BorderSide{},
				Left:   BorderSide{},
				Right:  BorderSide{},
			}
			surface := draw.NewBufferSurface(geometry.NewSize(10, 10))

			t.Run("Left", func(t *testing.T) {
				drawer.Left(style, surface)

				for x := 0; x < surface.Size().Width(); x++ {
					for y := 0; y < surface.Size().Height(); y++ {
						require.Equal(t, draw.Cell{}, surface.Get(geometry.NewVec2D(x, y)))
					}
				}
			})

			t.Run("Top", func(t *testing.T) {
				drawer.Top(style, surface)

				for x := 0; x < surface.Size().Width(); x++ {
					for y := 0; y < surface.Size().Height(); y++ {
						require.Equal(t, draw.Cell{}, surface.Get(geometry.NewVec2D(x, y)))
					}
				}
			})

			t.Run("Right", func(t *testing.T) {
				drawer.Right(style, surface)

				for x := 0; x < surface.Size().Width(); x++ {
					for y := 0; y < surface.Size().Height(); y++ {
						require.Equal(t, draw.Cell{}, surface.Get(geometry.NewVec2D(x, y)))
					}
				}
			})

			t.Run("Bottom", func(t *testing.T) {
				drawer.Bottom(style, surface)

				for x := 0; x < surface.Size().Width(); x++ {
					for y := 0; y < surface.Size().Height(); y++ {
						require.Equal(t, draw.Cell{}, surface.Get(geometry.NewVec2D(x, y)))
					}
				}
			})
		})
	})
}
