package styles

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/negrel/paon/pdk/events"
)

func BenchmarkStyle(b *testing.B) {
	b.Run("Style", func(b *testing.B) {
		benchmarkStyle(b, func() Style {
			return newStyle(events.NewTarget())
		})
	})

	b.Run("Theme", func(b *testing.B) {
		for i := 1; i < 64; i = i << 1 {
			b.Run(fmt.Sprintf("With-%d-Style", i), func(b *testing.B) {
				benchmarkStyle(b, func() Style {
					theme := NewTheme(nil)
					for j := 0; j < i; j++ {
						theme.AddStyle(NewWeighted(newStyle(events.NewTarget()), math.MinInt))
					}

					return theme
				})
			})
		}
	})
}

func benchmarkStyle(b *testing.B, newStyle func() Style) {
	b.Run("Bool", func(b *testing.B) {
		benchmarkBoolStyle(b, func() BoolStyle {
			return newStyle()
		})

		for i := 1; i <= 16; i = i << 1 {
			b.Run("Set-With-"+strconv.Itoa(i)+"-Listeners", func(b *testing.B) {
				benchmarkBoolStyleSet(b, func() BoolStyle {
					s := newStyle()
					for j := 0; j < i; j++ {
						s.AddEventListener(BoolChangedListener(func(bce BoolChangedEvent) {}))
					}
					return s
				})
			})
		}
	})

	b.Run("Color", func(b *testing.B) {
		benchmarkColorStyle(b, func() ColorStyle {
			return newStyle()
		})

		for i := 1; i <= 16; i = i << 1 {
			b.Run("Set-With-"+strconv.Itoa(i)+"-Listeners", func(b *testing.B) {
				benchmarkColorStyleSet(b, func() ColorStyle {
					s := newStyle()
					for j := 0; j < i; j++ {
						s.AddEventListener(ColorChangedListener(func(cce ColorChangedEvent) {}))
					}
					return s
				})
			})
		}
	})

	b.Run("Iface", func(b *testing.B) {
		benchmarkIfaceStyle(b, func() IfaceStyle {
			return newStyle()
		})

		for i := 1; i <= 16; i = i << 1 {
			b.Run("Set-With-"+strconv.Itoa(i)+"-Listeners", func(b *testing.B) {
				benchmarkIfaceStyleSet(b, func() IfaceStyle {
					s := newStyle()
					for j := 0; j < i; j++ {
						s.AddEventListener(IfaceChangedListener(func(ice IfaceChangedEvent) {}))
					}
					return s
				})
			})
		}
	})

	b.Run("Int", func(b *testing.B) {
		benchmarkIntStyle(b, func() IntStyle {
			return newStyle()
		})

		for i := 1; i <= 16; i = i << 1 {
			b.Run("Set-With-"+strconv.Itoa(i)+"-Listeners", func(b *testing.B) {
				benchmarkIntStyleSet(b, func() IntStyle {
					s := newStyle()
					for j := 0; j < i; j++ {
						s.AddEventListener(IfaceChangedListener(func(ice IfaceChangedEvent) {}))
					}
					return s
				})
			})
		}
	})

	b.Run("IntUnit", func(b *testing.B) {
		benchmarkIntUnitStyle(b, func() IntUnitStyle {
			return newStyle()
		})

		for i := 1; i <= 16; i = i << 1 {
			b.Run("Set-With-"+strconv.Itoa(i)+"-Listeners", func(b *testing.B) {
				benchmarkIntUnitStyleSet(b, func() IntUnitStyle {
					s := newStyle()
					for j := 0; j < i; j++ {
						s.AddEventListener(IntUnitChangedListener(func(ice IntUnitChangedEvent) {}))
					}
					return s
				})
			})
		}
	})
}
