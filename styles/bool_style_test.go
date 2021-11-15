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
		styleType:    reflect.TypeOf(boolStyle{}),
		defaultValue: (*property.Bool)(nil),

		new: func() reflect.Value {
			return reflect.ValueOf(newBoolStyle())
		},
		newID: func() reflect.Value {
			return reflect.ValueOf(property.NewBoolID("test-bool"))
		},
		newValue: func() reflect.Value {
			b := property.Bool(rand.Uint32()%2 == 0)
			return reflect.ValueOf(&b)
		},

		getMethodName: "Bool",
		setMethodName: "SetBool",
	}, styleDescription{
		styleType:    reflect.TypeOf((*style)(nil)),
		defaultValue: (*property.Bool)(nil),

		new: func() reflect.Value {
			return reflect.ValueOf(newStyle(events.NewTarget()))
		},
		newID: func() reflect.Value {
			return reflect.ValueOf(property.NewBoolID("test-bool"))
		},
		newValue: func() reflect.Value {
			b := property.Bool(rand.Uint32()%2 == 0)
			return reflect.ValueOf(&b)
		},

		getMethodName: "Bool",
		setMethodName: "SetBool",
	})
}

func BenchmarkBoolStyle(b *testing.B) {
	benchmarkBoolStyle(b, func() BoolStyle {
		return newBoolStyle()
	})
}

func benchmarkBoolStyle(b *testing.B, newStyle func() BoolStyle) {
	b.Run("Set", func(b *testing.B) {
		benchmarkBoolStyleSet(b, newStyle)
	})

	b.Run("Get", func(b *testing.B) {
		benchmarkBoolStyleGet(b, newStyle)
	})
}

func benchmarkBoolStyleSet(b *testing.B, newStyle func() BoolStyle) {
	for i := 16; i <= 2048; i = i << 1 {
		ids := make([]property.BoolID, i)
		bools := make([]property.Bool, i)
		for j := 0; j < i; j++ {
			ids[j] = property.NewBoolID(strconv.Itoa(j))
			bools[j] = property.Bool(rand.Uint32()%2 == 0)
		}

		b.Run(strconv.Itoa(i), func(b *testing.B) {
			style := newStyle()
			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				for k := 0; k < i; k++ {
					style.SetBool(ids[k], &bools[k])
				}
			}
		})
	}
}

func benchmarkBoolStyleGet(b *testing.B, newStyle func() BoolStyle) {
	for i := 16; i <= 2048; i = i << 1 {
		ids := make([]property.BoolID, i)
		for j := 0; j < i; j++ {
			ids[j] = property.NewBoolID(strconv.Itoa(j))
		}

		b.Run(strconv.Itoa(i), func(b *testing.B) {
			style := newStyle()
			for k := 0; k < i; k++ {
				b := property.True()
				if k%2 == 0 {
					b = property.False()
				}
				style.SetBool(ids[k], b)
			}

			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				for k := 0; k < i; k++ {
					c := style.Bool(ids[k])
					_ = c
				}
			}
		})
	}
}
