package widgets

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/internal/tree"
	"github.com/negrel/paon/pkg/pdk/displays"
	"github.com/negrel/paon/pkg/pdk/events"
	"github.com/negrel/paon/pkg/pdk/flows"
	"github.com/negrel/paon/pkg/pdk/render"
)

// Root define the Root node of the Widget tree.
type Root struct {
	*layout
	screen displays.Screen
	engine render.Engine
}

func NewRoot(screen displays.Screen, engine render.Engine, child Widget) *Root {
	r := &Root{
		screen: screen,
		engine: engine,
	}
	r.layout = newLayout("root", r)
	_, err := r.AppendChild(child)
	assert.Nil(err)

	r.screen.AddEventListener(events.ResizeListener(func(resize events.Resize) {
		engine.Enqueue(r.FirstChild())
	}))

	return r
}

// RootNode implements the tree.Node interface.
func (r *Root) RootNode() tree.ParentNode {
	return r
}

// Box implements the Widget interface.
func (r *Root) Box() flows.BoxModel {
	return flows.NewBox(r.screen.Bounds())
}

// Render implements the render.Renderable interface.
func (r *Root) Render() render.Patch {
	constraint := flows.MakeConstraint(
		geometry.Rectangle{}, r.screen.Bounds(),
		r.screen.Bounds().Size(), r.screen.Bounds().Size(),
	)

	childBox := r.FirstChild().Flow(constraint)
	canvas := r.screen.Canvas()
	ctx := canvas.NewContext(childBox.MarginBox())
	r.FirstChild().Draw(ctx)

	return canvas.Patch()
}
