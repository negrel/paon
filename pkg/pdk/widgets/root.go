package widgets

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/events"
	"github.com/negrel/paon/pkg/pdk/displays"
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
		layout: newLayout("root"),
		screen: screen,
		engine: engine,
	}
	err := r.AppendChild(child)
	assert.Nil(err)

	r.screen.AddEventListener(events.ResizeListener(func(resize events.Resize) {
		engine.Enqueue(r)
	}))

	return r
}

func (r *Root) Root() *Root {
	return r
}

// Box implements the Widget interface.
func (r *Root) Box() flows.BoxModel {
	return flows.NewBox(r.screen.Bounds())
}

func (r *Root) AppendChild(child Widget) error {
	err := r.layout.AppendChild(child)
	if err == nil {
		child.setParent(r)
	}

	return err
}

func (r *Root) InsertBefore(reference, child Widget) error {
	err := r.layout.InsertBefore(reference, child)
	if err == nil {
		child.setParent(r)
	}

	return err
}

// Render implements the render.Renderable interface.
func (r *Root) Render() render.Patch {
	constraint := flows.MakeConstraint(
		geometry.Rectangle{}, r.screen.Bounds(),
		r.screen.Bounds().Size(), r.screen.Bounds().Size(),
	)

	childBox := r.FirstChild().FlowAlgo()(constraint)
	canvas := r.screen.Canvas()
	ctx := canvas.NewContext(childBox.MarginBox())
	r.FirstChild().Drawer().Draw(ctx)

	return canvas.Patch()
}
