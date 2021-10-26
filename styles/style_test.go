package styles

import (
	"testing"

	"github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/styles/property"
	"github.com/stretchr/testify/assert"
)

func newTestStyle() Style {
	return New(events.NewTarget())
}

type styleTest struct {
	name      string
	functions []func(t *testing.T, s Style)
}

func TestStyle(t *testing.T) {
	styleTests := generateStyleTests()

	for _, methodTests := range styleTests {
		for _, tests := range methodTests.functions {
			t.Run(methodTests.name, func(t *testing.T) {
				tests(t, newTestStyle())
			})
		}
	}
}

func generateStyleTests() []styleTest {
	tests := []styleTest{
		{
			name: "Set",
			functions: []func(t *testing.T, s Style){
				testStyleSet,
				testStyleSetCustomProp,
			},
		},
		{
			name: "Get",
			functions: []func(t *testing.T, s Style){
				testStyleGet,
				testStyleGetCustomProp,
			},
		},
		{
			name: "Del",
			functions: []func(t *testing.T, s Style){
				testStyleDel,
			},
		},
	}

	return tests
}

func testStyleSet(t *testing.T, s Style) {
	p := s.Get(property.WidthID())
	assert.Nil(t, p)

	prop := property.NewProp(property.WidthID())
	assert.NotNil(t, prop)
	s.Set(prop)

	p = s.Get(property.WidthID())
	assert.Equal(t, prop, p)
}

func testStyleSetCustomProp(t *testing.T, s Style) {
	id := property.NewID("test-id")

	p := s.Get(id)
	assert.Nil(t, p)

	prop := property.NewProp(id)
	assert.NotNil(t, prop)
	s.Set(prop)

	p = s.Get(id)
	assert.Equal(t, prop, p)
}

func testStyleGet(t *testing.T, s Style) {
	prop := property.NewProp(property.PaddingBottomID())

	s.Set(prop)

	actual := s.Get(property.PaddingBottomID())

	assert.Equal(t, prop.ID(), actual.ID())
	assert.Equal(t, prop, actual)
}

func testStyleGetCustomProp(t *testing.T, s Style) {
	id := property.NewID("mock")
	prop := property.NewProp(id)

	s.Set(prop)

	actual := s.Get(id)

	assert.Equal(t, prop.ID(), actual.ID())
	assert.Equal(t, prop, actual)
}

func testStyleDel(t *testing.T, s Style) {
	prop := property.NewProp(property.WidthID())
	s.Set(prop)
	s.Del(property.WidthID())

	assert.Nil(t, s.Get(property.WidthID()))
}
