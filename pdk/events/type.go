package events

import (
	"fmt"

	"github.com/negrel/paon/pdk/id"
)

var typeRegistry = id.Registry{}
var typeMap = id.NewStringMap()

// Type is the type of an Event
type Type id.ID

func (t Type) name() string {
	return typeMap.Get(id.ID(t))
}

// String implements the fmt.Stringer interface.
func (t Type) String() string {
	return fmt.Sprintf("%v", t.name())
}

// NewType returns a new Type that can be used for custom events
// type.
func NewType(name string) Type {
	t := typeRegistry.New()
	typeMap.Set(t, name)

	return Type(t)
}
