package property

import "fmt"

//go:generate stringer -type=ID -trimprefix=ID
var _ fmt.Stringer = ID(0)

type ID int

const (
	IDWidth ID = iota + 1
	IDMinWidth
	IDMaxWidth

	IDHeight
	IDMinHeight
	IDMaxHeight

	IDMarginLeft
	IDMarginTop
	IDMarginRight
	IDMarginBottom

	IDPaddingLeft
	IDPaddingTop
	IDPaddingRight
	IDPaddingBottom

	IDBackgroundColor
)
