package render

import (
	"strconv"
	"testing"

	"github.com/negrel/paon/geometry"
)

func intPow(n, pow int) int {
	for i := 0; i < pow; i++ {
		n *= n
	}

	return n
}

func BenchmarkBufferSurface(b *testing.B) {
	b.Run("Resize", func(b *testing.B) {
		b.Run("Shrink", func(b *testing.B) {
			initialSize := geometry.NewSize(1<<10, 1<<10)
			for i := 6; i < 10; i++ {
				b.Run(strconv.Itoa(1<<i), func(b *testing.B) {
					benchmarkBufferSurfaceResize(b, initialSize, geometry.NewSize(1<<i, 1<<i))
				})
			}
		})

		b.Run("Grow", func(b *testing.B) {
			initialSize := geometry.NewSize(8, 8)
			for i := 6; i < 10; i++ {
				b.Run(strconv.Itoa(1<<i), func(b *testing.B) {
					benchmarkBufferSurfaceResize(b, initialSize, geometry.NewSize(1<<i, 1<<i))
				})
			}
		})
	})
}

func benchmarkBufferSurfaceResize(b *testing.B, initial geometry.Size, new geometry.Size) {
	bs := NewBufferSurface(initial)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bs.Resize(new)
	}

	_ = bs.Size()
}
