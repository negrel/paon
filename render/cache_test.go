package render

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("MarkDirty/IsForwarded", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		r := NewMockRenderable(ctrl)
		r.EXPECT().MarkDirty().Times(1)

		cache := NewCache(r)
		cache.MarkDirty()
	})

	t.Run("IsDirty/IsForwarded", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		r := NewMockRenderable(ctrl)
		r.EXPECT().IsDirty().Return(true).Times(1)

		cache := NewCache(r)
		require.True(t, cache.IsDirty())
	})

	t.Run("Layout", func(t *testing.T) {
		t.Run("DirtyRenderable", func(t *testing.T) {
			t.Run("SameConstraint/LayoutCallForwarded", func(t *testing.T) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()
				r := NewMockRenderable(ctrl)

				cache := NewCache(r)

				r.EXPECT().IsDirty().Return(true)
				r.EXPECT().Layout(layout.Constraint{}).Return(geometry.Size{Width: 10, Height: 0})
				size := cache.Layout(layout.Constraint{})

				require.Equal(t, geometry.Size{Width: 10, Height: 0}, size)
			})

			t.Run("DifferentConstraint/LayoutCallForwarded", func(t *testing.T) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()
				r := NewMockRenderable(ctrl)

				cache := NewCache(r)
				// Cached constraint must be zero on creation.
				require.Equal(t, layout.Constraint{}, cache.cachedConstraint)

				co := layout.Constraint{
					MinSize:    geometry.Size{},
					MaxSize:    geometry.Size{Width: 100, Height: 0},
					ParentSize: geometry.Size{},
					RootSize:   geometry.Size{},
				}

				r.EXPECT().IsDirty().Return(true)
				r.EXPECT().Layout(co).Return(geometry.Size{Width: 10, Height: 0})
				size := cache.Layout(co)

				require.Equal(t, geometry.Size{Width: 10, Height: 0}, size)
			})
		})

		t.Run("CleanRenderable", func(t *testing.T) {
			t.Run("SameConstraint/CachedSizeReturned", func(t *testing.T) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()
				r := NewMockRenderable(ctrl)

				cache := NewCache(r)

				r.EXPECT().IsDirty().Return(false)
				r.EXPECT().Layout(layout.Constraint{}).Times(0)
				size := cache.Layout(layout.Constraint{})

				require.Equal(t, geometry.Size{}, size)
			})

			t.Run("DifferentConstraint/LayoutCallForwarded", func(t *testing.T) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()
				r := NewMockRenderable(ctrl)

				cache := NewCache(r)
				// Cached constraint must be zero on creation.
				require.Equal(t, layout.Constraint{}, cache.cachedConstraint)

				co := layout.Constraint{
					MinSize:    geometry.Size{},
					MaxSize:    geometry.Size{Width: 100, Height: 0},
					ParentSize: geometry.Size{},
					RootSize:   geometry.Size{},
				}

				r.EXPECT().IsDirty().Return(false)
				r.EXPECT().Layout(co).Return(geometry.Size{Width: 10, Height: 0})
				size := cache.Layout(co)

				require.Equal(t, geometry.Size{Width: 10, Height: 0}, size)
			})
		})
	})

	t.Run("Draw", func(t *testing.T) {
		t.Run("DrawCallForwarded", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			r := NewMockRenderable(ctrl)

			cache := NewCache(r)

			r.EXPECT().Draw(nil).Times(1)
			cache.Draw(nil)
		})
	})
}
