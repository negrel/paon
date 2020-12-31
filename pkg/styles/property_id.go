package styles

import "fmt"

//go:generate stringer -type=PropID -trimprefix=PropID
var _ fmt.Stringer = PropID(0)

type PropID int

const (
	PropIDWidth PropID = iota
	PropIDMinWidth
	PropIDMaxWidth

	PropIDHeight
	PropIDMinHeight
	PropIDMaxHeight

	PropIDMarginLeft
	PropIDMarginTop
	PropIDMarginRight
	PropIDMarginBottom

	PropIDPaddingLeft
	PropIDPaddingTop
	PropIDPaddingRight
	PropIDPaddingBottom
)
