package widgets

import (
	"github.com/negrel/paon/draw"
)

// Enhancer define a widget that can have a single child.
type Enhancer interface {
	Widget

	Widget() Widget
	// SwapWidget swaps single child.
	SwapWidget(Widget) (Widget, error)
}

// BaseEnhancer define a minimal/incomplete and unopiniated wrapper
// implementation that should be embedded into more complex implementation.
type BaseEnhancer struct {
	BaseWidget
	ChildLayout
}

// NewBaseEnhancer returns a new base enhancer.
func NewBaseEnhancer(embedder Widget) BaseEnhancer {
	return BaseEnhancer{
		BaseWidget: NewBaseWidget(embedder),
	}

}

// SwapWidget swaps single child widget and returns old widget.
func (be *BaseEnhancer) SwapWidget(w Widget) (Widget, error) {
	old := be.Widget()
	if old != nil {
		be.node.RemoveChild(nodeOrNil(old))
	}
	err := be.node.AppendChild(nodeOrNil(w))
	return old, err
}

// Widget returns single child widget.
func (be *BaseEnhancer) Widget() Widget {
	return widgetOrNil(be.node.FirstChild())
}

// Draw implements draw.Drawer.
func (be *BaseEnhancer) Draw(srf draw.Surface) {
	widget := be.Widget()
	if widget == nil {
		return
	}

	subsrf := draw.NewSubSurface(srf, be.ChildLayout.Bounds)
	widget.Draw(subsrf)
}
