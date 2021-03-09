package events

import (
	"fmt"
	"github.com/negrel/paon/pkg/pdk/id"
)

// Type is the type of an Event
type Type id.ID

func (t Type) name() string {
	return typeName.Get(id.ID(t))
}

func (t Type) String() string {
	return fmt.Sprintf("%v (%d)", t.name(), t)
}

var typeName = id.NewMap()

func MakeType(name string) Type {
	t := id.Make()
	typeName.Set(t, name)

	return Type(t)
}

// List of existing event type.
var (
	_TypeError       = MakeType("error")
	_TypeUnsupported = MakeType("unsupported")
	_TypeInterrupt   = MakeType("interrupt")
	_TypeClick       = MakeType("click")
	_TypeMouseMove   = MakeType("mouse-move")
	_TypeKeyboard    = MakeType("keyboard")
	_TypeResize      = MakeType("resize")
	_TypeWheel       = MakeType("wheel")
)

func TypeError() Type {
	return _TypeError
}

func TypeUnsupported() Type {
	return _TypeUnsupported
}

func TypeInterrupt() Type {
	return _TypeInterrupt
}

func TypeClick() Type {
	return _TypeClick
}

func TypeMouseMove() Type {
	return _TypeMouseMove
}

func TypeKeyboard() Type {
	return _TypeKeyboard
}

func TypeResize() Type {
	return _TypeResize
}

func TypeWheel() Type {
	return _TypeWheel
}
