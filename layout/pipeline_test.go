package layout

import (
	"testing"

	"github.com/negrel/paon/geometry"
	"github.com/stretchr/testify/require"
)

func TestPipeline(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		pipeline := NewPipeline(LayoutFunc(func(co Constraint) geometry.Size {
			return co.MaxSize
		}))

		size := pipeline.Layout(Constraint{
			MinSize:    geometry.Size{},
			MaxSize:    geometry.NewSize(10, 10),
			ParentSize: geometry.Size{},
			RootSize:   geometry.Size{},
		})

		require.Equal(t, geometry.NewSize(10, 10), size)
	})

	t.Run("Double", func(t *testing.T) {
		pipeline := NewPipeline(LayoutFunc(func(co Constraint) geometry.Size {
			return co.MaxSize
		}))

		pipeline.Pipe(func(co Constraint, next Layout) geometry.Size {
			size := next.Layout(co)
			return geometry.NewSize(size.Width()*2, size.Height()*2)
		})

		size := pipeline.Layout(Constraint{
			MinSize:    geometry.Size{},
			MaxSize:    geometry.NewSize(10, 10),
			ParentSize: geometry.Size{},
			RootSize:   geometry.Size{},
		})

		require.Equal(t, geometry.NewSize(20, 20), size)
	})

	t.Run("Triple", func(t *testing.T) {
		pipeline := NewPipeline(LayoutFunc(func(co Constraint) geometry.Size {
			return co.MaxSize
		}))

		pipeline.Pipe(func(co Constraint, next Layout) geometry.Size {
			size := next.Layout(co)
			return geometry.NewSize(size.Width()*2, size.Height()*2)
		})

		pipeline.Pipe(func(co Constraint, next Layout) geometry.Size {
			size := next.Layout(co)
			return geometry.NewSize(size.Width()+5, size.Height()+5)
		})

		size := pipeline.Layout(Constraint{
			MinSize:    geometry.Size{},
			MaxSize:    geometry.NewSize(10, 10),
			ParentSize: geometry.Size{},
			RootSize:   geometry.Size{},
		})

		require.Equal(t, geometry.NewSize(25, 25), size)
	})
}
