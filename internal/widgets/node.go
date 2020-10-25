package widgets

import (
	"fmt"
	"github.com/google/uuid"

	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/internal/style"
)

var _ Widget = &Node{}

// Node define is a minimal leaf node in the Widget tree.
type Node struct {
	uuid   uuid.UUID
	name   string
	parent Layout
	style  *style.Node
}

func (n *Node) slot() int {
	panic("implement me")
}

// NewNodeWidget return a new Node object to embed in custom
// widget.
func NewNodeWidget(name string) *Node {
	n := &Node{
		name:  name,
		uuid:  uuid.New(),
		style: style.Unset(),
	}

	log.Infoln("creating", n)

	return n
}

func (n *Node) Name() string {
	return n.name
}

func (n *Node) UUID() uuid.UUID {
	return n.uuid
}

func (n *Node) String() string {
	return fmt.Sprintf("%v-%v", n.name, n.uuid.String())
}

func (n *Node) Parent() Layout {
	return n.parent
}

func (n *Node) adoptedBy(parent Layout) {
	n.parent = parent
	// TODO update style tree.
}

func (n *Node) Style() *style.Node {
	return n.style
}
