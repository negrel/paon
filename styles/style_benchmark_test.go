package styles

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/negrel/paon/styles/property"
)

func BenchmarkMain(b *testing.B) {
	b.Run("Style", benchmarkStyle)
}

func benchmarkStyle(b *testing.B) {
	b.Run("Set", benchmarksStyleSet)
	b.Run("Get", benchmarksStyleGet)
}

func benchmarksStyleSet(b *testing.B) {
	for i := 8; i < 1024; i *= 2 {
		b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
			benchmarkStyleSetInt(b, i)
		})
	}
}

func benchmarkStyleSetInt(b *testing.B, n int) {
	ids := make([]property.IntID, n)
	props := make([]property.Int, n)

	for i := 0; i < n; i++ {
		ids = append(ids, property.NewIntID(strconv.Itoa(i)))
		props = append(props, property.NewInt(i))
	}
	style := newTestStyle()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			style.SetInt(ids[j], &props[n])
		}
	}
}

func benchmarksStyleGet(b *testing.B) {
	for i := 8; i < 1024; i *= 2 {
		b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
			benchmarkStyleGet(b, i)
		})
	}
}

func benchmarkStyleGet(b *testing.B, n int) {
	ids := []property.IntID{}

	for i := 0; i < n; i++ {
		ids = append(ids, property.NewIntID(strconv.Itoa(i)))
	}
	style := newTestStyle()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, id := range ids {
			_ = style.Int(id)
		}
	}
}
