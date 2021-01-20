package events

import (
	"fmt"
	"github.com/negrel/paon/internal/idmap"
	"strings"
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
	TypeMouseMove
	TypeKeyboard
	TypeResize
	TypeWheel

	// Used for custom property created using the pdk
	unusedType
)

func (t Type) name() string {
	return typeName.Get(int32(t))
}

func (t Type) String() string {
	return fmt.Sprintf("%v (%d)", t.name(), t)
}

var typeName = idmap.New(int(unusedType))

func init() {
	typeName.Set(int32(TypeError), "error")
	typeName.Set(int32(TypeUnsupported), "unsupported")
	typeName.Set(int32(TypeInterrupt), "interrupt")
	typeName.Set(int32(TypeClick), "click")
	typeName.Set(int32(TypeMouseMove), "mouse-move")
	typeName.Set(int32(TypeKeyboard), "keyboard")
	typeName.Set(int32(TypeResize), "resize")
	typeName.Set(int32(TypeWheel), "wheel")
}

var eventTypeCounter int32 = int32(unusedType - 1)

func MakeType(name string) Type {
	t := atomic.AddInt32(&eventTypeCounter, 1)
	typeName.Set(t, strings.ToLower(name))
	return t
}
