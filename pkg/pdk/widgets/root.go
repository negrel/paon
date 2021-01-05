package widgets

import (
	"github.com/negrel/paon/internal/tree"
)

type rLayout = Layout

type Root interface {
	tree.Root
	Layout
}

var _ tree.Root = &root{}

// root define the root node of the Widget tree.
type root struct {
	rLayout
	children Widget
}

func NewRoot(child Widget) *root {
	r := &root{
		rLayout: newLayout(tree.NewRoot(child)),
	}

	return r
}
