package widgets

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/tree"
)

func nodeOrNil(w Widget) *tree.Node[Widget] {
	if w == nil {
		return nil
	}

	return w.Node()
}

func widgetOrNil(n *tree.Node[Widget]) Widget {
	if n == nil {
		return nil
	}

	return n.Unwrap()
}

// BaseLayout define a minimal/incomplete widget implementation that should be
// embedded into more complex implementation.
type BaseLayout struct {
	BaseWidget
	ChildrenLayout ChildrenLayout
}

// NewBaseLayout returns a new base layout.
func NewBaseLayout(embedder Widget) BaseLayout {
	bl := BaseLayout{
		BaseWidget:     NewBaseWidget(embedder),
		ChildrenLayout: ChildrenLayout{},
	}

	return bl
}

func (bl *BaseLayout) DispatchEvent(ev events.Event) {
	// Trigger events listeners.
	bl.BaseWidget.DispatchEvent(ev)

	// Forward to child widgets.
	if data, isPointerData := ev.Data.(events.PointerEventData); isPointerData {
		for i := 0; i < bl.ChildrenLayout.Size(); i++ {
			childLayout := bl.ChildrenLayout.Get(i)
			if childLayout.Bounds.Contains(data.RelativePosition()) {
				ev.Data = data.WithPositionRelativeToOrigin(childLayout.Bounds.Origin)
				childLayout.Widget.DispatchEvent(ev)
			}
		}
	} else {
		child := bl.FirstChild()
		for child != nil {
			child.DispatchEvent(ev)
			child = child.NextSibling()
		}
	}
}

// Draw implements draw.Drawer.
func (bl *BaseLayout) Draw(srf draw.Surface) {
	bl.ChildrenLayout.Draw(srf)
}

// FirstChild returns first child of this layout.
func (bl *BaseLayout) FirstChild() Widget {
	return widgetOrNil(bl.node.FirstChild())
}

// LastChild returns last child of this layout.
func (bl *BaseLayout) LastChild() Widget {
	return widgetOrNil(bl.node.LastChild())
}

func (bl *BaseLayout) AppendChild(w Widget) error {
	return bl.node.AppendChild(nodeOrNil(w))
}

func (bl *BaseLayout) InsertBefore(newChild, reference Widget) error {
	return bl.node.InsertBefore(nodeOrNil(newChild), nodeOrNil(reference))
}

func (bl *BaseLayout) RemoveChild(child Widget) error {
	return bl.node.RemoveChild(nodeOrNil(child))
}

func (bl *BaseLayout) IsAncestorOf(other Widget) bool {
	return bl.node.IsAncestorOf(nodeOrNil(other))
}

func (bl *BaseLayout) IsDescendantOf(other Widget) bool {
	return bl.node.IsDescendantOf(nodeOrNil(other))
}
