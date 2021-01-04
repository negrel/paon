package property

import (
	"fmt"
	"sync/atomic"
)

//go:generate stringer -type=ID -trimprefix=ID
var _ fmt.Stringer = ID(0)

type ID int32

const (
	IDDisplay ID = iota + 1

	IDWidth
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

	// Used for custom property created using the pdk
	unusedID
)

var propIdCounter int32 = int32(unusedID - 1)

func MakeID() ID {
	return ID(atomic.AddInt32(&propIdCounter, 1))
}
