package events

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
)

// Type is the type of an Event
type Type int32

// List of existing event type.
const (
	TypeError Type = iota - 2
	TypeUnsupported
	TypeInterrupt
	TypeClick
	TypeKeyboard
	TypeResize
	TypeWheel

	// Used for custom property created using the pdk
	unusedType
)

var eventTypeCounter int32 = int32(unusedType - 1)

func (t Type) name() string {
	return typeName.get(t)
}

func (t Type) String() string {
	return fmt.Sprintf("%v (%d)", t.name(), t)
}

type typeMap struct {
	sync.Mutex
	m map[Type]string
}

func (tm *typeMap) set(p Type, name string) {
	tm.Lock()
	defer tm.Unlock()
	tm.m[p] = name
}

func (tm *typeMap) get(p Type) string {
	tm.Lock()
	defer tm.Unlock()
	return tm.m[p]
}

var typeName = &typeMap{
	Mutex: sync.Mutex{},

	m: map[Type]string{
		TypeError:       "error",
		TypeUnsupported: "unsupported",
		TypeInterrupt:   "interrupt",
		TypeClick:       "click",
		TypeKeyboard:    "keyboard",
		TypeResize:      "resize",
		TypeWheel:       "wheel",
	},
}

func MakeType(name string) Type {
	t := Type(atomic.AddInt32(&eventTypeCounter, 1))
	typeName.set(t, strings.ToLower(name))
	return t
}
