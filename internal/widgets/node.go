package widgets

import (
	"fmt"
	"github.com/google/uuid"

	"github.com/negrel/paon/internal/style"
)

var _ Widget = &Node{}

// Node define is a minimal leaf node in the Widget tree.
type Node struct {
	uuid   uuid.UUID
	name   string
	parent Layout
	style  style.Node
}

// NewNodeWidget return a new Node object to embed in custom
// widget.
func NewNodeWidget(name string) *Node {
	return &Node{
		name: name,
		uuid: uuid.New(),
	}
}

func (c *Node) Name() string {
	return c.name
}

func (c *Node) UUID() uuid.UUID {
	return c.uuid
}

func (c *Node) String() string {
	return fmt.Sprintf("%v-%v", c.name, c.uuid.String())
}

func (c *Node) Parent() Layout {
	return c.parent
}

func (c *Node) setParent(parent Layout) {
	c.parent = parent
}
