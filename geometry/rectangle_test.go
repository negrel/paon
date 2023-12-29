package geometry

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRectange(t *testing.T) {
	t.Run("Getter", func(t *testing.T) {
		r := Rectangle{Vec2D{3, 6}, Size{10, 20}}

		require.Equal(t, Size{10, 20}, r.Size())
		require.Equal(t, Size{10, 20}, r.RectSize)

		require.Equal(t, Vec2D{3, 6}, r.TopLeft())
		require.Equal(t, Vec2D{3, 6}, r.Origin)

		require.Equal(t, Vec2D{13, 6}, r.TopRight())
		require.Equal(t, Vec2D{13, 26}, r.BottomRight())
		require.Equal(t, Vec2D{3, 26}, r.BottomLeft())

		require.Equal(t, 3, r.Left())
		require.Equal(t, 6, r.Top())
		require.Equal(t, 13, r.Right())
		require.Equal(t, 26, r.Bottom())
	})

	t.Run("Contains", func(t *testing.T) {
		r := Rectangle{Vec2D{3, 6}, Size{10, 20}}

		require.True(t, r.Contains(r.TopLeft()))
		require.False(t, r.Contains(r.TopRight()))
		require.False(t, r.Contains(r.BottomRight()))
		require.False(t, r.Contains(r.BottomLeft()))

		require.True(t, r.Contains(Vec2D{5, 7}))
	})

	t.Run("Grow", func(t *testing.T) {
		r1 := Rectangle{
			Origin:   Vec2D{},
			RectSize: Size{100, 100},
		}

		t.Run("Left", func(t *testing.T) {
			r2 := r1.GrowLeft(3)
			require.Equal(t, Vec2D{-3, 0}, r2.Origin)
			require.Equal(t, -3, r2.Left())
			require.Equal(t, 0, r2.Top())
			require.Equal(t, 100, r2.Right())
			require.Equal(t, 100, r2.Bottom())
			require.NotEqual(t, r1, r2)
		})

		t.Run("Top", func(t *testing.T) {
			r2 := r1.GrowTop(3)
			require.Equal(t, Vec2D{0, -3}, r2.Origin)
			require.Equal(t, 0, r2.Left())
			require.Equal(t, -3, r2.Top())
			require.Equal(t, 100, r2.Right())
			require.Equal(t, 100, r2.Bottom())
			require.NotEqual(t, r1, r2)
		})

		t.Run("Right", func(t *testing.T) {
			r2 := r1.GrowRight(3)
			require.Equal(t, Vec2D{0, 0}, r2.Origin)
			require.Equal(t, 0, r2.Left())
			require.Equal(t, 0, r2.Top())
			require.Equal(t, 103, r2.Right())
			require.Equal(t, 100, r2.Bottom())
			require.NotEqual(t, r1, r2)
		})

		t.Run("Bottom", func(t *testing.T) {
			r2 := r1.GrowBottom(3)
			require.Equal(t, Vec2D{0, 0}, r2.Origin)
			require.Equal(t, 0, r2.Left())
			require.Equal(t, 0, r2.Top())
			require.Equal(t, 100, r2.Right())
			require.Equal(t, 103, r2.Bottom())
			require.NotEqual(t, r1, r2)
		})
	})

	t.Run("Shrink", func(t *testing.T) {
		r1 := Rectangle{
			Origin:   Vec2D{},
			RectSize: Size{100, 100},
		}

		t.Run("Left", func(t *testing.T) {
			r2 := r1.ShrinkLeft(3)
			require.Equal(t, Vec2D{3, 0}, r2.Origin)
			require.Equal(t, 3, r2.Left())
			require.Equal(t, 0, r2.Top())
			require.Equal(t, 100, r2.Right())
			require.Equal(t, 100, r2.Bottom())
			require.NotEqual(t, r1, r2)
		})

		t.Run("Top", func(t *testing.T) {
			r2 := r1.ShrinkTop(3)
			require.Equal(t, Vec2D{0, 3}, r2.Origin)
			require.Equal(t, 0, r2.Left())
			require.Equal(t, 3, r2.Top())
			require.Equal(t, 100, r2.Right())
			require.Equal(t, 100, r2.Bottom())
			require.NotEqual(t, r1, r2)
		})

		t.Run("Right", func(t *testing.T) {
			r2 := r1.ShrinkRight(3)
			require.Equal(t, Vec2D{0, 0}, r2.Origin)
			require.Equal(t, 0, r2.Left())
			require.Equal(t, 0, r2.Top())
			require.Equal(t, 97, r2.Right())
			require.Equal(t, 100, r2.Bottom())
			require.NotEqual(t, r1, r2)
		})

		t.Run("Bottom", func(t *testing.T) {
			r2 := r1.ShrinkBottom(3)
			require.Equal(t, Vec2D{0, 0}, r2.Origin)
			require.Equal(t, 0, r2.Left())
			require.Equal(t, 0, r2.Top())
			require.Equal(t, 100, r2.Right())
			require.Equal(t, 97, r2.Bottom())
			require.NotEqual(t, r1, r2)
		})
	})

	t.Run("MoveBy", func(t *testing.T) {
		r1 := Rectangle{
			Origin:   Vec2D{100, 100},
			RectSize: Size{100, 100},
		}

		r2 := r1.MoveBy(Vec2D{1, 2})
		require.Equal(t, Vec2D{101, 102}, r2.Origin)
		require.Equal(t, r1.RectSize, r2.RectSize)

		require.NotEqual(t, r1, r2)
	})
}
