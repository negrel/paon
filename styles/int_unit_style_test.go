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
		styleType:    reflect.TypeOf(intUnitStyle{}),
		defaultValue: (*property.IntUnit)(nil),

		new: func() reflect.Value {
			return reflect.ValueOf(newIntUnitStyle())
		},
		newID: func() reflect.Value {
			return reflect.ValueOf(property.NewIntUnitID("test-int-unit"))
		},
		newValue: func() reflect.Value {
			result := property.NewIntUnit(rand.Int(), property.CellUnit)
			return reflect.ValueOf(&result)
		},
		getMethodName: "IntUnit",
		setMethodName: "SetIntUnit",
	}, styleDescription{
		styleType:    reflect.TypeOf((*style)(nil)),
		defaultValue: (*property.IntUnit)(nil),

		new: func() reflect.Value {
			return reflect.ValueOf(newStyle(events.NewTarget()))
		},
		newID: func() reflect.Value {
			return reflect.ValueOf(property.NewIntUnitID("test-int-unit"))
		},
		newValue: func() reflect.Value {
			result := property.NewIntUnit(rand.Int(), property.CellUnit)
			return reflect.ValueOf(&result)
		},
		getMethodName: "IntUnit",
		setMethodName: "SetIntUnit",
	})
}

func BenchmarkIntUnitStyle(b *testing.B) {
	benchmarkIntUnitStyle(b, func() IntUnitStyle {
		return newIntUnitStyle()
	})
}

func benchmarkIntUnitStyle(b *testing.B, newStyle func() IntUnitStyle) {
	b.Run("Set", func(b *testing.B) {
		benchmarkIntUnitStyleSet(b, newStyle)
	})
	b.Run("Get", func(b *testing.B) {
		benchmarkIntUnitStyleGet(b, newStyle)
	})
}

func benchmarkIntUnitStyleSet(b *testing.B, newStyle func() IntUnitStyle) {
	for i := 16; i <= 2048; i = i << 1 {
		iter := property.NewIntUnit(i, property.CellUnit)

		ids := make([]property.IntUnitID, i)
		for j := 0; j < i; j++ {
			ids[j] = property.NewIntUnitID(strconv.Itoa(j))
		}

		b.Run(strconv.Itoa(i), func(b *testing.B) {
			style := newStyle()
			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				for k := 0; k < i; k++ {
					style.SetIntUnit(ids[k], &iter)
				}
			}
		})
	}
}

func benchmarkIntUnitStyleGet(b *testing.B, newStyle func() IntUnitStyle) {
	for i := 16; i <= 2048; i = i << 1 {
		iter := property.NewIntUnit(i, property.CellUnit)

		ids := make([]property.IntUnitID, i)
		for j := 0; j < i; j++ {
			ids[j] = property.NewIntUnitID(strconv.Itoa(j))
		}

		b.Run(strconv.Itoa(i), func(b *testing.B) {
			style := newStyle()
			for k := 0; k < i; k++ {
				style.SetIntUnit(ids[k], &iter)
			}

			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				for k := 0; k < i; k++ {
					c := style.IntUnit(ids[k])
					_ = c
				}
			}
		})
	}
}
