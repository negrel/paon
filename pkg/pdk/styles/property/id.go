package property

import (
	"fmt"
	"sync"
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
	return idName.get(id)
}

func (id ID) String() string {
	return fmt.Sprintf("%v (%d)", id.name(), id)
}

type idMap struct {
	sync.Mutex
	m map[ID]string
}

func (im *idMap) set(i ID, name string) {
	im.Lock()
	defer im.Unlock()
	im.m[i] = name
}

func (im *idMap) get(i ID) string {
	im.Lock()
	defer im.Unlock()
	return im.m[i]
}

var idName = &idMap{
	Mutex: sync.Mutex{},
	m: map[ID]string{
		IDDisplay:         "display",
		IDWidth:           "width",
		IDMinWidth:        "min-width",
		IDMaxWidth:        "max-width",
		IDHeight:          "height",
		IDMinHeight:       "min-height",
		IDMaxHeight:       "max-height",
		IDMarginLeft:      "margin-left",
		IDMarginTop:       "margin-top",
		IDMarginRight:     "margin-right",
		IDMarginBottom:    "margin-bottom",
		IDPaddingLeft:     "padding-left",
		IDPaddingTop:      "padding-top",
		IDPaddingRight:    "padding-right",
		IDPaddingBottom:   "padding-bottom",
		IDBackgroundColor: "background-color",
	},
}

var propIdCounter int32 = int32(unusedID - 1)

func MakeID(name string) ID {
	id := ID(atomic.AddInt32(&propIdCounter, 1))
	idName.set(id, name)
	return id
}
