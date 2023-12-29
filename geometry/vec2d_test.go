package geometry

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVec2D(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		v1 := Vec2D{1, 2}
		v2 := Vec2D{3, 4}

		result := v1.Add(v2)

		require.Equal(t, 4, result.X)
		require.Equal(t, 6, result.Y)

		// v1 & v2 remains unchanged.
		require.Equal(t, 1, v1.X)
		require.Equal(t, 2, v1.Y)
		require.Equal(t, 3, v2.X)
		require.Equal(t, 4, v2.Y)
	})

	t.Run("Sub", func(t *testing.T) {
		v1 := Vec2D{1, 2}
		v2 := Vec2D{3, 4}

		result := v1.Sub(v2)

		require.Equal(t, -2, result.X)
		require.Equal(t, -2, result.Y)

		// v1 & v2 remains unchanged.
		require.Equal(t, 1, v1.X)
		require.Equal(t, 2, v1.Y)
		require.Equal(t, 3, v2.X)
		require.Equal(t, 4, v2.Y)
	})
}
