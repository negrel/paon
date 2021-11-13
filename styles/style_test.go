package styles

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/negrel/paon/styles/property"
	"github.com/stretchr/testify/assert"
)

type styleDescription struct {
	styleType    reflect.Type
	defaultValue interface{}

	new      func() reflect.Value
	newID    func() reflect.Value
	newValue func() reflect.Value

	getMethodName string
	setMethodName string
}

var stylesDesc = []styleDescription{}

func TestStyles(t *testing.T) {
	for _, styleDesc := range stylesDesc {
		if styleDesc.styleType.Kind() == reflect.Ptr {
			styleDesc.styleType = styleDesc.styleType.Elem()
		}

		t.Run(styleDesc.styleType.Name(), func(t *testing.T) {
			testStyle(t, styleDesc)
		})
	}
}

func testStyle(t *testing.T, styleDesc styleDescription) {
	t.Run(styleDesc.setMethodName, func(t *testing.T) {
		styleProp := styleDesc.newID()
		style := styleDesc.new()
		set := style.MethodByName(styleDesc.setMethodName)
		in := make([]reflect.Value, 2)
		in[0] = styleProp
		in[1] = styleDesc.newValue()

		t.Run("Property", func(t *testing.T) {
			results := set.Call(in)
			assert.Len(t, results, 0)

			get := style.MethodByName(styleDesc.getMethodName)
			results = get.Call(in[:1])
			assert.Equal(t, in[1].Interface(), results[0].Interface())
		})

		t.Run("NewProperty", func(t *testing.T) {
			in[0] = styleDesc.newID()
			assert.Panics(t, func() {
				set.Call(in)
			})
		})
	})

	t.Run(styleDesc.getMethodName, func(t *testing.T) {
		styleProp := styleDesc.newID()
		style := styleDesc.new()

		get := style.MethodByName(styleDesc.getMethodName)

		t.Run("DefaultValue", func(t *testing.T) {
			result := get.Call([]reflect.Value{styleProp})[0]
			assert.Equal(t, styleDesc.defaultValue, result.Interface())
		})

		t.Run("NewProperty", func(t *testing.T) {
			assert.Panics(t, func() {
				get.Call([]reflect.Value{
					styleDesc.newID(),
				})
			})
		})
	})
}

func TestStyleEvents(t *testing.T) {
	t.Run("Color", testStyleEventsColor)
	t.Run("Iface", testStyleEventsIface)
	t.Run("Int", testStyleEventsInt)
	t.Run("IntUnit", testStyleEventsIntUnit)
}

func testStyleEventsColor(t *testing.T) {
	var old, new *property.Color
	var eventTriggered bool
	id := property.NewColorID("test-color")

	style := New(nil)
	style.AddEventListener(ColorChangedListener(func(cce ColorChangedEvent) {
		eventTriggered = true
		assert.Equal(t, id, cce.ColorID)
		assert.Equal(t, old, cce.Old)
		assert.Equal(t, new, cce.New)
	}))

	new = &property.ColorBlueViolet
	style.SetColor(id, new)

	old = new
	new = &property.ColorPink
	style.SetColor(id, new)

	old = new
	new = nil
	style.SetColor(id, new)

	assert.True(t, eventTriggered)
}

func testStyleEventsIface(t *testing.T) {
	var old, new interface{}
	var eventTriggered bool
	id := property.NewIfaceID("test-iface")

	style := New(nil)
	style.AddEventListener(IfaceChangedListener(func(ice IfaceChangedEvent) {
		eventTriggered = true
		assert.Equal(t, id, ice.IfaceID)
		assert.Equal(t, old, ice.Old)
		assert.Equal(t, new, ice.New)
	}))

	new = time.Now()
	style.SetIface(id, new)

	old = new
	new = rand.Int()
	style.SetIface(id, new)

	old = new
	new = nil
	style.SetIface(id, new)

	assert.True(t, eventTriggered)
}

func randomPropInt() *property.Int {
	a := property.NewInt(rand.Int())
	return &a
}

func testStyleEventsInt(t *testing.T) {
	var old, new *property.Int
	var eventTriggered bool
	id := property.NewIntID("test-int")

	style := New(nil)
	style.AddEventListener(IntChangedListener(func(ice IntChangedEvent) {
		eventTriggered = true
		assert.Equal(t, id, ice.IntID)
		assert.Equal(t, old, ice.Old)
		assert.Equal(t, new, ice.New)
	}))

	new = randomPropInt()
	style.SetInt(id, new)

	old = new
	new = randomPropInt()
	style.SetInt(id, new)

	old = new
	new = nil
	style.SetInt(id, new)

	assert.True(t, eventTriggered)
}

func randomPropIntUnit() *property.IntUnit {
	a := property.NewIntUnit(rand.Int(), property.CellUnit)
	return &a
}

func testStyleEventsIntUnit(t *testing.T) {
	var old, new *property.IntUnit
	var eventTriggered bool
	id := property.NewIntUnitID("test-int-unit")

	style := New(nil)
	style.AddEventListener(IntUnitChangedListener(func(iuce IntUnitChangedEvent) {
		eventTriggered = true
		assert.Equal(t, id, iuce.IntID)
		assert.Equal(t, old, iuce.Old)
		assert.Equal(t, new, iuce.New)
	}))

	new = randomPropIntUnit()
	style.SetIntUnit(id, new)

	old = new
	new = randomPropIntUnit()
	style.SetIntUnit(id, new)

	old = new
	new = nil
	style.SetIntUnit(id, new)

	assert.True(t, eventTriggered)
}
