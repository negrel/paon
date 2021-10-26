package styles

import (
	"fmt"
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
			benchmarkStyleSet(b, i)
		})
		b.Run(fmt.Sprintf("custom-prop %d", i), func(b *testing.B) {
			benchmarkStyleSetCustomProp(b, i)
		})
	}
}

func benchmarkStyleSet(b *testing.B, n int) {
	style := newTestStyle()
	props := make([]property.Property, n)

	builtInPropsCount := int(property.LastID())

	for j := 0; j < n; j++ {
		id := property.ID(j % builtInPropsCount)
		props = append(props, property.NewProp(id))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			style.Set(props[n])
		}
	}
}

func benchmarkStyleSetCustomProp(b *testing.B, n int) {
	style := newTestStyle()
	props := make([]property.Property, n)

	for j := 0; j < n; j++ {
		props = append(props, property.NewProp(property.NewID("test")))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			style.Set(props[n])
		}
	}
}

func benchmarksStyleGet(b *testing.B) {
	for i := 8; i < 1024; i *= 2 {
		b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
			benchmarkStyleGet(b, i)
		})
		b.Run(fmt.Sprintf("custom-prop %d", i), func(b *testing.B) {
			benchmarkStyleGetCustomProps(b, i)
		})
	}
}

func benchmarkStyleGet(b *testing.B, n int) {
	style := newTestStyle()
	ids := []property.ID{}

	last := int(property.LastID())

	for i := 0; i < n; i++ {
		ids = append(ids, property.ID(i%last))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, id := range ids {
			_ = style.Get(id)
		}
	}
}

func benchmarkStyleGetCustomProps(b *testing.B, n int) {
	style := newTestStyle()
	ids := []property.ID{}

	for i := 0; i < n; i++ {
		ids = append(ids, property.NewID(fmt.Sprintf("mock-%d", i)))
		style.Set(property.NewInt(ids[i], i))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, id := range ids {
			_ = style.Get(id)
		}
	}
}
