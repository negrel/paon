package events

import (
	"fmt"

	"github.com/negrel/paon/pdk/id"
)

var typeRegistry = id.Registry{}
var typeMap = map[Type]string{}

// Type is the type of an Event
type Type id.ID

func (t Type) name() string {
	return typeMap[t]
}

// String implements the fmt.Stringer interface.
func (t Type) String() string {
	return fmt.Sprintf("%v", t.name())
}

// NewType returns a new Type that can be used for custom events
// type. All events type must be created before application starts.
func NewType(name string) Type {
	t := Type(typeRegistry.New())
	typeMap[t] = name

	return t
}
