package property

import (
	"fmt"

	"github.com/negrel/paon/pdk/id"
)

// ID define a unique int32 number.
type ID id.ID

var idName = id.NewMap()

func (i ID) name() string {
	return idName.Get(id.ID(i))
}

func (i ID) String() string {
	return fmt.Sprintf("%v (%d)", i.name(), i)
}

func NewID(name string) ID {
	i := id.New()
	idName.Set(i, name)
	return ID(i)
}

var (
	_IDFlow ID = NewID("flow")

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

func Flow() ID {
	return _IDFlow
}

func Width() ID {
	return _IDWidth
}

func MinWidth() ID {
	return _IDMinWidth
}

func MaxWidth() ID {
	return _IDMaxWidth
}

func Height() ID {
	return _IDHeight
}

func MinHeight() ID {
	return _IDMinHeight
}

func MaxHeight() ID {
	return _IDMaxHeight
}

func MarginLeft() ID {
	return _IDMarginLeft
}

func MarginTop() ID {
	return _IDMarginTop
}

func MarginRight() ID {
	return _IDMarginRight
}

func MarginBottom() ID {
	return _IDMarginBottom
}

func BorderLeft() ID {
	return _IDBorderLeft
}

func BorderTop() ID {
	return _IDBorderTop
}

func BorderRight() ID {
	return _IDBorderRight
}

func BorderBottom() ID {
	return _IDBorderBottom
}

func BorderCharset() ID {
	return _IDBorderCharSet
}

func BorderColor() ID {
	return _IDBorderColor
}

func PaddingLeft() ID {
	return _IDPaddingLeft
}

func PaddingTop() ID {
	return _IDPaddingTop
}

func PaddingRight() ID {
	return _IDPaddingRight
}

func PaddingBottom() ID {
	return _IDPaddingBottom
}

func BackgroundColor() ID {
	return _IDBackgroundColor
}

func ForegroundColor() ID {
	return _IDForegroundColor
}
