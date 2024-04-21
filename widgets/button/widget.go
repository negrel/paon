package button

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events/mouse"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/render"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/tree"
	"github.com/negrel/paon/widgets"
	"github.com/negrel/paon/widgets/span"
)

type Option func(*Widget)

// WithStyle return an option that sets button widget style.
func WithStyle(style widgets.Style) Option {
	return func(w *Widget) {
		w.Cache.Unwrap().Renderable.style.InnerStyle = style
	}
}

func OnClick(handler func(event mouse.ClickEvent)) Option {
	return func(w *Widget) {
		w.AddEventListener(mouse.ClickListener(handler))
	}
}

type renderable struct {
	render.VoidRenderable
	text  string
	style widgets.InheritStyle
}

func newRenderable(nodeAccessor tree.NodeAccessor, text string) render.Cache[*styles.Renderable[*renderable]] {
	r := &renderable{
		VoidRenderable: render.NewVoidRenderable(nodeAccessor),
		text:           text,
		style: widgets.InheritStyle{
			NodeAccessor: nodeAccessor,
			InnerStyle:   widgets.Style{}.Margin(0).Padding(0).Border(styles.BorderSide{Size: 0}),
		},
	}

	return render.NewCache(&styles.Renderable[*renderable]{
		Renderable: r,
		Styled:     r,
	})
}

func (r *renderable) Style() styles.Style {
	return r.style
}

func (r *renderable) Layout(co layout.Constraint) geometry.Size {
	return span.Layout(co, r.text)
}

func (r *renderable) Draw(surface draw.Surface) {
	r.VoidRenderable.Draw(surface)

	style := r.style.Compute()
	span.Draw(surface, r.text, style.CellStyle)
}

// Widget define a clickable widget button.
type Widget struct {
	*widgets.PanicWidget
	render.Cache[*styles.Renderable[*renderable]]
}

// New returns a new button widget configured with the given options.
func New(text string, options ...Option) *Widget {
	w := &Widget{}

	w.PanicWidget = widgets.NewPanicWidget(w)
	w.Cache = newRenderable(w, text)

	w.AddEventListener(mouse.PressListener(func(event mouse.Event) {
		renderable := w.Cache.Unwrap().Renderable
		style := renderable.style
		renderable.style.InnerStyle = style.InnerStyle.Reverse(true)
		w.Cache.MarkDirty()
	}))

	w.AddEventListener(mouse.ClickListener(func(event mouse.ClickEvent) {
		renderable := w.Cache.Unwrap().Renderable
		style := renderable.style
		renderable.style.InnerStyle = style.InnerStyle.Reverse(false)
		w.Cache.MarkDirty()
	}))

	for _, applyOption := range options {
		applyOption(w)
	}

	return w
}

func (w *Widget) Renderable() render.Renderable {
	return &w.Cache
}
