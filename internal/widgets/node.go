package widgets

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/negrel/debuggo/pkg/log"

	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/internal/style"
	"github.com/negrel/paon/internal/utils"
)

var _ Widget = &Node{}

// Node define is a minimal leaf node in the Widget tree.
type Node struct {
	uuid   uuid.UUID
	name   string
	parent Layout
	style  *style.Node
}

// NewNodeWidget return a new Node object to embed in custom
// widget.
func NewNodeWidget(name string) *Node {
	n := &Node{
		name:  name,
		uuid:  uuid.New(),
		style: style.Unset(),
	}

	log.Infoln("creating", n, "widget")

	return n
}

// Name implements the Widget interface.
func (n *Node) Name() string {
	return n.name
}

// UUID implements the Widget interface.
func (n *Node) UUID() uuid.UUID {
	return n.uuid
}

// String implements the fmt.Stringer interface.
func (n *Node) String() string {
	return fmt.Sprintf("%v-%v", n.name, n.uuid.String())
}

// Parent implements the Widget interface.
func (n *Node) Parent() Layout {
	return n.parent
}

func (n *Node) adoptedBy(parent Layout) {
	n.parent = parent
	// TODO update style tree.
}

// Style implements the Widget interface.
func (n *Node) Style() *style.Node {
	return n.style
}

func (n *Node) slot() int {
	panic("implement me")
}

// Render implements the render.Renderable interface.
func (n *Node) Render(rect utils.Rectangle) render.Patch {
	panic("implement me")
}
