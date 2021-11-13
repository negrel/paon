package styles

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"

	"github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/styles/property"
)

func init() {
	stylesDesc = append(stylesDesc, styleDescription{
		styleType:    reflect.TypeOf(ifaceStyle{}),
		defaultValue: nil,

		new: func() reflect.Value {
			return reflect.ValueOf(newIfaceStyle())
		},
		newID: func() reflect.Value {
			return reflect.ValueOf(property.NewIfaceID("test-iface"))
		},
		newValue: func() reflect.Value {
			return reflect.ValueOf(interface{}(rand.Int()))
		},

		getMethodName: "Iface",
		setMethodName: "SetIface",
	}, styleDescription{
		styleType:    reflect.TypeOf((*style)(nil)),
		defaultValue: nil,

		new: func() reflect.Value {
			return reflect.ValueOf(newStyle(events.NewTarget()))
		},
		newID: func() reflect.Value {
			return reflect.ValueOf(property.NewIfaceID("test-iface"))
		},
		newValue: func() reflect.Value {
			return reflect.ValueOf(interface{}(rand.Int()))
		},

		getMethodName: "Iface",
		setMethodName: "SetIface",
	})
}

func BenchmarkIfaceStyle(b *testing.B) {
	benchmarkIfaceStyle(b, func() IfaceStyle {
		return newIfaceStyle()
	})
}

func benchmarkIfaceStyle(b *testing.B, newStyle func() IfaceStyle) {
	b.Run("Set", func(b *testing.B) {
		benchmarkIfaceStyleSet(b, newStyle)
	})

	b.Run("Get", func(b *testing.B) {
		benchmarkIfaceStyleGet(b, newStyle)
	})
}

func benchmarkIfaceStyleSet(b *testing.B, newStyle func() IfaceStyle) {
	for i := 16; i <= 2048; i = i << 1 {
		ids := make([]property.IfaceID, i)
		ifaces := make([]interface{}, i)
		for j := 0; j < i; j++ {
			ids[j] = property.NewIfaceID(strconv.Itoa(j))
			ifaces[j] = rand.Uint32()
		}

		b.Run(strconv.Itoa(i), func(b *testing.B) {
			style := newStyle()
			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				for k := 0; k < i; k++ {
					style.SetIface(ids[k], &ifaces[k])
				}
			}
		})
	}
}

func benchmarkIfaceStyleGet(b *testing.B, newStyle func() IfaceStyle) {
	for i := 16; i <= 2048; i = i << 1 {
		ids := make([]property.IfaceID, i)
		for j := 0; j < i; j++ {
			ids[j] = property.NewIfaceID(strconv.Itoa(j))
		}

		b.Run(strconv.Itoa(i), func(b *testing.B) {
			style := newStyle()
			for k := 0; k < i; k++ {
				style.SetIface(ids[k], rand.Int())
			}

			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				for k := 0; k < i; k++ {
					c := style.Iface(ids[k])
					_ = c
				}
			}
		})
	}
}
