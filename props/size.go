package props

import (
	"github.com/negrel/paon/styles/property"
	"github.com/negrel/paon/styles/value"
)

// Width returns a property that define the width of an element.
func Width(value value.Unit) property.Unit {
	return property.NewUnit(property.WidthID(), value)
}

// MinWidth returns a property that define the minimum width of an element.
func MinWidth(value value.Unit) property.Unit {
	return property.NewUnit(property.MinWidthID(), value)
}

// MaxWidth returns a property that define the maximum width of an element.
func MaxWidth(value value.Unit) property.Unit {
	return property.NewUnit(property.MaxWidthID(), value)
}

// Height returns a property that define the height of an element.
func Height(value value.Unit) property.Unit {
	return property.NewUnit(property.HeightID(), value)
}

// MinHeight returns a property that define the minimum width of an element.
func MinHeight(value value.Unit) property.Unit {
	return property.NewUnit(property.MinHeightID(), value)
}

// MaxHeight returns a property that define the maximum width of an element.
func MaxHeight(value value.Unit) property.Unit {
	return property.NewUnit(property.MaxHeightID(), value)
}
