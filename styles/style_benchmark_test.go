package styles

import (
	"strconv"
	"testing"

	"github.com/negrel/paon/pdk/events"
)

func BenchmarkStyle(b *testing.B) {
	b.Run("Color", func(b *testing.B) {
		benchmarkColorStyle(b, func() ColorStyle {
			return newStyle(events.NewTarget())
		})

		for i := 4; i <= 16; i = i << 1 {
			b.Run("Set-With-"+strconv.Itoa(i)+"-Listeners", func(b *testing.B) {
				benchmarkColorStyleSet(b, func() ColorStyle {
					s := newStyle(events.NewTarget())
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
			return newStyle(events.NewTarget())
		})

		for i := 4; i <= 16; i = i << 1 {
			b.Run("Set-With-"+strconv.Itoa(i)+"-Listeners", func(b *testing.B) {
				benchmarkIfaceStyleSet(b, func() IfaceStyle {
					s := newStyle(events.NewTarget())
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
			return newStyle(events.NewTarget())
		})

		for i := 4; i <= 16; i = i << 1 {
			b.Run("Set-With-"+strconv.Itoa(i)+"-Listeners", func(b *testing.B) {
				benchmarkIntStyleSet(b, func() IntStyle {
					s := newStyle(events.NewTarget())
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
			return newStyle(events.NewTarget())
		})

		for i := 4; i <= 16; i = i << 1 {
			b.Run("Set-With-"+strconv.Itoa(i)+"-Listeners", func(b *testing.B) {
				benchmarkIntUnitStyleSet(b, func() IntUnitStyle {
					s := newStyle(events.NewTarget())
					for j := 0; j < i; j++ {
						s.AddEventListener(IntUnitChangedListener(func(ice IntUnitChangedEvent) {}))
					}
					return s
				})
			})
		}
	})
}
