package styles

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/styles/property"
)

func init() {
	rand.Seed(time.Now().Unix())

	stylesDesc = append(stylesDesc, styleDescription{
		styleType:    reflect.TypeOf(colorStyle{}),
		defaultValue: (*property.Color)(nil),

		new: func() reflect.Value {
			return reflect.ValueOf(newColorStyle())
		},
		newID: func() reflect.Value {
			return reflect.ValueOf(property.NewColorID("test-color"))
		},
		newValue: func() reflect.Value {
			c := property.ColorFromHexa(rand.Uint32())
			return reflect.ValueOf(&c)
		},

		getMethodName: "Color",
		setMethodName: "SetColor",
	}, styleDescription{
		styleType:    reflect.TypeOf((*style)(nil)),
		defaultValue: (*property.Color)(nil),

		new: func() reflect.Value {
			return reflect.ValueOf(newStyle(events.NewTarget()))
		},
		newID: func() reflect.Value {
			return reflect.ValueOf(property.NewColorID("test-color"))
		},
		newValue: func() reflect.Value {
			c := property.ColorFromHexa(rand.Uint32())
			return reflect.ValueOf(&c)
		},

		getMethodName: "Color",
		setMethodName: "SetColor",
	})
}

func BenchmarkColorStyle(b *testing.B) {
	benchmarkColorStyle(b, func() ColorStyle {
		return newColorStyle()
	})
}

func benchmarkColorStyle(b *testing.B, newStyle func() ColorStyle) {
	b.Run("Set", func(b *testing.B) {
		benchmarkColorStyleSet(b, newStyle)
	})

	b.Run("Get", func(b *testing.B) {
		benchmarkColorStyleGet(b, newStyle)
	})
}

func benchmarkColorStyleSet(b *testing.B, newStyle func() ColorStyle) {
	for i := 16; i <= 2048; i = i << 1 {
		ids := make([]property.ColorID, i)
		colors := make([]property.Color, i)
		for j := 0; j < i; j++ {
			ids[j] = property.NewColorID(strconv.Itoa(j))
			colors[j] = property.ColorFromHex(rand.Uint32())
		}

		b.Run(strconv.Itoa(i), func(b *testing.B) {
			style := newStyle()
			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				for k := 0; k < i; k++ {
					style.SetColor(ids[k], &colors[k])
				}
			}
		})
	}
}

func benchmarkColorStyleGet(b *testing.B, newStyle func() ColorStyle) {
	for i := 16; i <= 2048; i = i << 1 {
		ids := make([]property.ColorID, i)
		for j := 0; j < i; j++ {
			ids[j] = property.NewColorID(strconv.Itoa(j))
		}

		b.Run(strconv.Itoa(i), func(b *testing.B) {
			style := newStyle()
			for k := 0; k < i; k++ {
				style.SetColor(ids[k], &property.ColorWhite)
			}

			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				for k := 0; k < i; k++ {
					c := style.Color(ids[k])
					_ = c
				}
			}
		})
	}
}
