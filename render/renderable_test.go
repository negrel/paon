package render

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/tree"
	"github.com/stretchr/testify/require"
)

func TestVoidRenderable(t *testing.T) {
	t.Run("NewVoidRenderable", func(t *testing.T) {
		t.Run("IsDirty", func(t *testing.T) {
			nodeAccessor := tree.NewNode(nil)
			vr := NewVoidRenderable(nodeAccessor)
			require.True(t, vr.IsDirty(), "void renderable is clean")
		})
	})

	t.Run("Layout", func(t *testing.T) {
		t.Run("ZeroSize", func(t *testing.T) {
			nodeAccessor := tree.NewNode(nil)
			vr := NewVoidRenderable(nodeAccessor)
			size := vr.Layout(layout.Constraint{})

			require.Equal(t, geometry.Size{}, size)
		})

		t.Run("RemainsDirty", func(t *testing.T) {
			nodeAccessor := tree.NewNode(nil)
			vr := NewVoidRenderable(nodeAccessor)
			_ = vr.Layout(layout.Constraint{})

			require.True(t, vr.IsDirty(), "void renderable is clean")
		})
	})

	t.Run("Draw", func(t *testing.T) {
		t.Run("SurfaceRemainsUntouched", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			surface := NewMockSurface(ctrl)
			defer ctrl.Finish()
			// No expected call.

			nodeAccessor := tree.NewNode(nil)
			vr := NewVoidRenderable(nodeAccessor)
			vr.Draw(surface)
		})

		t.Run("CleanAfterDraw", func(t *testing.T) {
			nodeAccessor := tree.NewNode(nil)
			vr := NewVoidRenderable(nodeAccessor)
			vr.Draw(nil)
			require.False(t, vr.IsDirty(), "void renderable is dirty")
		})
	})

	t.Run("MarkDirty", func(t *testing.T) {
		t.Run("DirtyRemainsDirty", func(t *testing.T) {
			nodeAccessor := tree.NewNode(nil)
			vr := NewVoidRenderable(nodeAccessor)
			vr.MarkDirty()

			require.True(t, vr.IsDirty(), "void renderable is clean")
		})

		t.Run("CleanBecomesDirty", func(t *testing.T) {
			nodeAccessor := tree.NewNode(nil)
			vr := NewVoidRenderable(nodeAccessor)

			// Make clean
			vr.Draw(nil)
			require.False(t, vr.IsDirty(), "void renderable is dirty")

			vr.MarkDirty()
			require.True(t, vr.IsDirty(), "void renderable is clean")
		})
	})
}
