package property

import (
	"fmt"

	"github.com/negrel/paon/pdk/id"
)

// ID define a unique uint32 number.
type ID id.ID

var idName = id.NewMap()
var propRegistry = id.Registry{}

func (i ID) name() string {
	return idName.Get(id.ID(i))
}

func (i ID) String() string {
	return fmt.Sprintf("%v (%d)", i.name(), i)
}

// NewID generates a unique property ID with the given name.
func NewID(name string) ID {
	i := propRegistry.New()
	idName.Set(i, name)
	return ID(i)
}

var (
	_IDWidth    = NewID("width")
	_IDMinWidth = NewID("min-width")
	_IDMaxWidth = NewID("max-width")

	_IDHeight    = NewID("height")
	_IDMinHeight = NewID("min-height")
	_IDMaxHeight = NewID("max-height")

	_IDMarginLeft   = NewID("margin-left")
	_IDMarginTop    = NewID("margin-top")
	_IDMarginRight  = NewID("margin-right")
	_IDMarginBottom = NewID("margin-bottom")

	_IDBorderLeft    = NewID("border-left")
	_IDBorderTop     = NewID("border-top")
	_IDBorderRight   = NewID("border-right")
	_IDBorderBottom  = NewID("border-bottom")
	_IDBorderColor   = NewID("border-color")
	_IDBorderCharSet = NewID("border-charset")

	_IDPaddingLeft   = NewID("padding-left")
	_IDPaddingTop    = NewID("padding-top")
	_IDPaddingRight  = NewID("padding-right")
	_IDPaddingBottom = NewID("padding-bottom")

	_IDBackgroundColor = NewID("background-color")
	_IDForegroundColor = NewID("foreground-color")
)

// IsCustomPropID returns true if the given ID is not one of the built-in property ID.
func IsCustomPropID(id ID) bool {
	return id > LastID()
}

// LastID returns the last property ID parts this package.
func LastID() ID {
	return _IDForegroundColor
}

// WidthID returns the ID of the Width property.
func WidthID() ID {
	return _IDWidth
}

// MinWidthID returns the ID of the MinWidth property.
func MinWidthID() ID {
	return _IDMinWidth
}

// MaxWidthID returns the ID of the MaxWidth property.
func MaxWidthID() ID {
	return _IDMaxWidth
}

// HeightID returns the ID of the Height property.
func HeightID() ID {
	return _IDHeight
}

// MinHeightID returns the ID of the MinHeight property.
func MinHeightID() ID {
	return _IDMinHeight
}

// MaxHeightID returns the ID of the MaxHeight property.
func MaxHeightID() ID {
	return _IDMaxHeight
}

// MarginLeftID returns the ID of the MarginLeft property.
func MarginLeftID() ID {
	return _IDMarginLeft
}

// MarginTopID returns the ID of the MarginTop property.
func MarginTopID() ID {
	return _IDMarginTop
}

// MarginRightID returns the ID of the MarginRight property.
func MarginRightID() ID {
	return _IDMarginRight
}

// MarginBottomID returns the ID of the MarginBottom property.
func MarginBottomID() ID {
	return _IDMarginBottom
}

// BorderLeftID returns the ID of the BorderLeft property.
func BorderLeftID() ID {
	return _IDBorderLeft
}

// BorderTopID returns the ID of the BorderTop property.
func BorderTopID() ID {
	return _IDBorderTop
}

// BorderRightID returns the ID of the BorderRight property.
func BorderRightID() ID {
	return _IDBorderRight
}

// BorderBottomID returns the ID of the BorderBottom property.
func BorderBottomID() ID {
	return _IDBorderBottom
}

// BorderCharsetID returns the ID of the BorderCharset property.
func BorderCharsetID() ID {
	return _IDBorderCharSet
}

// BorderColorID returns the ID of the BorderColor property.
func BorderColorID() ID {
	return _IDBorderColor
}

// PaddingLeftID returns the ID of the PaddingLeft property.
func PaddingLeftID() ID {
	return _IDPaddingLeft
}

// PaddingTopID returns the ID of the PaddingTop property.
func PaddingTopID() ID {
	return _IDPaddingTop
}

// PaddingRightID returns the ID of the PaddingRight property.
func PaddingRightID() ID {
	return _IDPaddingRight
}

// PaddingBottomID returns the ID of the PaddingBottom property.
func PaddingBottomID() ID {
	return _IDPaddingBottom
}

// BackgroundColorID returns the ID of the BackgroundColor property.
func BackgroundColorID() ID {
	return _IDBackgroundColor
}

// ForegroundColorID returns the ID of the ForegroundColor property.
func ForegroundColorID() ID {
	return _IDForegroundColor
}
