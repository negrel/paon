package widgets

import (
	"errors"

	"github.com/negrel/paon/events"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/internal/metrics"
	"github.com/negrel/paon/pdk/draw"
	pdkevents "github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/pdk/layout"
	"github.com/negrel/paon/pdk/render"
	"github.com/negrel/paon/pdk/tree"
)

// Root define the root of a widget tree.
type Root struct {
	*BaseWidget

	dst draw.Surface

	child     Widget
	childRect geometry.Rectangle
	layer     render.Layer
}

var _ Layout = &Root{}

// NewRoot returns a new Widget that can be used as a root.
func NewRoot(target pdkevents.Target, dst draw.Surface) *Root {
	root := &Root{
		layer: render.Layer{
			BufferSurface: render.NewBufferSurface(geometry.NewSize(80, 24)),
		},
		dst: dst,
	}

	// Update root layer size on resize.
	target.AddEventListener(events.ResizeListener(func(r events.Resize) {
		root.layer.BufferSurface = root.layer.BufferSurface.Resize(r.New)
	}))

	root.BaseWidget = newBaseWidget(
		Wrap(root),
		Target(target),
		Renderable(root),
		NodeConstructor(
			func(data interface{}) tree.Node {
				return tree.NewRoot(data)
			},
		),
	)
	root.layer.Node = root.BaseWidget.node

	return root
}

// AppendChild implements the Layout interface.
func (r *Root) AppendChild(child Widget) error {
	if r.child == nil {
		r.SetChild(child)
		return nil
	}

	return errors.New("can't append child, root can only have one child")
}

// InsertBefore implements the Layout interface.
func (r *Root) InsertBefore(reference, newChild Widget) error {
	if reference != nil {
		return errors.New("can't insert child, the given reference must be nil on a root node")
	}

	return r.AppendChild(newChild)
}

// RemoveChild implements the Layout interface.
func (r *Root) RemoveChild(child Widget) error {
	if r.child.IsSame(child) {
		r.SetChild(nil)
		return nil
	}

	return errors.New("can't remove child, the widget is not a child of the root")
}

// FirstChild implements the Layout interface.
func (r Root) FirstChild() Widget {
	return r.child
}

// LastChild implements the Layout interface.
func (r Root) LastChild() Widget {
	return r.child
}

// Root returns itself to implements the Widget interface.
func (r *Root) Root() *Root {
	return r
}

// SetChild sets the direct child of the root.
// If a child is already present, it is unmounted.
func (r *Root) SetChild(child Widget) {
	if r.child != nil {
		r.BaseWidget.Node().RemoveChild(r.child.Node())
	}

	r.child = child
	if child != nil {
		r.BaseWidget.Node().AppendChild(child.Node())
	}
}

// Layer implements the render.Renderable interface.
func (r *Root) Layer() *render.Layer {
	return &r.layer
}

// Render implements the render.Renderable interface.
func (r Root) Render(ctx render.Context) {
	metrics.StartRenderTimer()
	defer metrics.StopRenderTimer()

	r.child.Render(ctx)
}

func (r Root) layoutConstraint() layout.Constraint {
	return layout.Constraint{
		MinSize:    geometry.Size{},
		MaxSize:    r.layer.Size(),
		RootSize:   r.layer.Size(),
		ParentSize: r.layer.Size(),
	}
}

// PerformRender starts the rendering of the Root layer.
func (r Root) PerformRender() {
	ctx := render.Context{
		Layer:      r.layer,
		Layout:     r,
		Constraint: r.layoutConstraint(),
	}

	r.Render(ctx)

	for x := 0; x < r.dst.Size().Width(); x++ {
		for y := 0; y < r.dst.Size().Height(); y++ {
			pos := geometry.NewVec2D(x, y)
			r.dst.Set(pos, r.layer.Get(pos))
		}
	}
}

// Layout implements the layout.Layout interface.
func (r Root) Layout(sized geometry.Sized) geometry.Rectangle {
	size := sized.Size()

	// Resizing to the same size clear the layer
	r.layer.BufferSurface.Resize(r.layer.BufferSurface.Size())

	return geometry.Rectangle{
		Min: geometry.Vec2D{},
		Max: geometry.NewVec2D(size.Width(), size.Height()),
	}
}
