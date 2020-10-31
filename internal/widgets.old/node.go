package widgets

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/negrel/debuggo/pkg/log"

	"github.com/negrel/paon/internal/events"
	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/internal/style"
	"github.com/negrel/paon/internal/utils"
)

var _ Widget = &Node{}

// Node define is a minimal leaf node in the Widget tree.
type Node struct {
	events.Target
	utils.Rectangle

	uuid   uuid.UUID
	name   string
	parent Layout
	window Widget
	rect   utils.Rectangle

	style *style.Node
}

// NewNodeWidget return a new Node object to embed in custom
// widget.
func NewNodeWidget(name string, options ...Option) *Node {
	n := &Node{
		Target:    events.MakeTarget(),
		Rectangle: utils.Rect(0, 0, 0, 0),

		name:  name,
		uuid:  uuid.New(),
		style: style.Unset(),
	}

	for _, option := range options {
		option(n)
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

// Window implements the Widget interface.
func (n *Node) Window() Widget {
	return n.window
}

func (n *Node) isAttached() bool {
	return n.window != nil
}

func (n *Node) adoptedBy(parent Layout) {
	n.parent = parent
	if n.parent.isAttached() {
		n.window = n.parent.Window()
	}

	// TODO update style tree.
}

// Style implements the Widget interface.
func (n *Node) Style() *style.Node {
	return n.style
}

// Render implements the render.Renderable interface.
func (n *Node) Render(_ utils.Rectangle) render.Patch {
	panic("implement me")
}

// DispatchEvent implements the events.EventTarget interface.
func (n *Node) DispatchEvent(event events.Event) {
	if ce, ok := event.(events.ClickEvent); ok {
		if !n.Contain(ce.Position) {
			return
		}
	}

	n.Target.DispatchEvent(event)
}
