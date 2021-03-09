package property

import (
	"fmt"
	"github.com/negrel/paon/pkg/pdk/id"
)

type ID id.ID

var idName = id.NewMap()

func (i ID) name() string {
	return idName.Get(id.ID(i))
}

func (i ID) String() string {
	return fmt.Sprintf("%v (%d)", i.name(), i)
}

func MakeID(name string) ID {
	i := id.Make()
	idName.Set(i, name)
	return ID(i)
}

var (
	_IDFlow ID = MakeID("flow")

	_IDWidth    = MakeID("width")
	_IDMinWidth = MakeID("min-width")
	_IDMaxWidth = MakeID("max-width")

	_IDHeight    = MakeID("height")
	_IDMinHeight = MakeID("min-height")
	_IDMaxHeight = MakeID("max-height")

	_IDMarginLeft   = MakeID("margin-left")
	_IDMarginTop    = MakeID("margin-top")
	_IDMarginRight  = MakeID("margin-right")
	_IDMarginBottom = MakeID("margin-bottom")

	_IDBorderLeft    = MakeID("border-left")
	_IDBorderTop     = MakeID("border-top")
	_IDBorderRight   = MakeID("border-right")
	_IDBorderBottom  = MakeID("border-bottom")
	_IDBorderColor   = MakeID("border-color")
	_IDBorderCharSet = MakeID("border-charset")

	_IDPaddingLeft   = MakeID("padding-left")
	_IDPaddingTop    = MakeID("padding-top")
	_IDPaddingRight  = MakeID("padding-right")
	_IDPaddingBottom = MakeID("padding-bottom")

	_IDBackgroundColor = MakeID("background-color")
	_IDForegroundColor = MakeID("foreground-color")
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
