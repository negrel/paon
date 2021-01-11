package property

import (
	"fmt"
	"github.com/negrel/paon/internal/idmap"
	"sync/atomic"
)

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

func (id ID) name() string {
	return idName.Get(int32(id))
}

func (id ID) String() string {
	return fmt.Sprintf("%v (%d)", id.name(), id)
}

var idName = idmap.New(int(unusedID))

var propIdCounter int32 = int32(unusedID - 1)

func MakeID(name string) ID {
	id := atomic.AddInt32(&propIdCounter, 1)
	idName.Set(id, name)
	return ID(id)
}
