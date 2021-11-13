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
		styleType:    reflect.TypeOf(intStyle{}),
		defaultValue: (*property.Int)(nil),

		new: func() reflect.Value {
			return reflect.ValueOf(newIntStyle())
		},
		newID: func() reflect.Value {
			return reflect.ValueOf(property.NewIntID("test-int"))
		},
		newValue: func() reflect.Value {
			result := property.NewInt(rand.Int())
			return reflect.ValueOf(&result)
		},
		getMethodName: "Int",
		setMethodName: "SetInt",
	}, styleDescription{
		styleType:    reflect.TypeOf((*style)(nil)),
		defaultValue: (*property.Int)(nil),

		new: func() reflect.Value {
			return reflect.ValueOf(newStyle(events.NewTarget()))
		},
		newID: func() reflect.Value {
			return reflect.ValueOf(property.NewIntID("test-int"))
		},
		newValue: func() reflect.Value {
			result := property.NewInt(rand.Int())
			return reflect.ValueOf(&result)
		},
		getMethodName: "Int",
		setMethodName: "SetInt",
	})
}

func BenchmarkIntStyle(b *testing.B) {
	benchmarkIntStyle(b, func() IntStyle {
		return newIntStyle()
	})
}

func benchmarkIntStyle(b *testing.B, newStyle func() IntStyle) {
	b.Run("Set", func(b *testing.B) {
		benchmarkIntStyleSet(b, newStyle)
	})
	b.Run("Get", func(b *testing.B) {
		benchmarkIntStyleGet(b, newStyle)
	})
}

func benchmarkIntStyleSet(b *testing.B, newStyle func() IntStyle) {
	for i := 16; i <= 2048; i = i << 1 {
		iter := property.NewInt(i)

		ids := make([]property.IntID, i)
		for j := 0; j < i; j++ {
			ids[j] = property.NewIntID(strconv.Itoa(j))
		}

		b.Run(strconv.Itoa(i), func(b *testing.B) {
			style := newStyle()
			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				for k := 0; k < i; k++ {
					style.SetInt(ids[k], &iter)
				}
			}
		})
	}
}

func benchmarkIntStyleGet(b *testing.B, newStyle func() IntStyle) {
	for i := 16; i <= 2048; i = i << 1 {
		iter := property.NewInt(i)

		ids := make([]property.IntID, i)
		for j := 0; j < i; j++ {
			ids[j] = property.NewIntID(strconv.Itoa(j))
		}

		b.Run(strconv.Itoa(i), func(b *testing.B) {
			style := newStyle()
			for k := 0; k < i; k++ {
				style.SetInt(ids[k], &iter)
			}

			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				for k := 0; k < i; k++ {
					c := style.Int(ids[k])
					_ = c
				}
			}
		})
	}
}
