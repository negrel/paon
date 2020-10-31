package widgets

import (
	"fmt"

	"github.com/google/uuid"
)

type Widget interface {
	fmt.Stringer

	ID() uuid.UUID
	Name() string

	isSame(Widget) bool
}

var _ Widget = &widget{}

type widget struct {
	name string
	id   uuid.UUID
}

func newWidget(name string) *widget {
	return &widget{
		name: name,
		id:   uuid.New(),
	}
}

func (w *widget) String() string {
	return fmt.Sprintf("%v-%v", w.name, w.id)
}

func (w *widget) ID() uuid.UUID {
	return w.id
}

func (w *widget) Name() string {
	return w.name
}

func (w *widget) isSame(other Widget) bool {
	return w.ID() == other.ID()
}
