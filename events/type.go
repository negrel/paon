package events

import (
	"fmt"

	"github.com/negrel/paon/pdk/id"
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

func NewType(name string) Type {
	t := id.New()
	typeName.Set(t, name)

	return Type(t)
}
