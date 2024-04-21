package input

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/render"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/tree"
	"github.com/negrel/paon/widgets"
)

type Option func(*Widget)

// WithStyle define input internal styles.Style.
func WithStyle(style widgets.Style) Option {
	return func(w *Widget) {
		w.Cache.Unwrap().Renderable.style.InnerStyle = style
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
	return Layout(co, r.text)
}

func (r *renderable) Draw(surface draw.Surface) {
	r.VoidRenderable.Draw(surface)

	Draw(surface, r.text, r.style.Compute().CellStyle)
}

type Widget struct {
	*widgets.PanicWidget
	render.Cache[*styles.Renderable[*renderable]]
}

func New(text string, options ...Option) *Widget {
	w := &Widget{}
	w.PanicWidget = widgets.NewPanicWidget(w)
	w.Cache = newRenderable(w, text)

	for _, applyOption := range options {
		applyOption(w)
	}

	return w
}

func (w *Widget) Renderable() render.Renderable {
	return &w.Cache
}

func (w Widget) SetText(txt string) {
	renderable := w.Cache.Unwrap()
	renderable.Renderable.text = txt
	w.Cache.MarkDirty()
}

func (w Widget) Text() string {
	renderable := w.Cache.Unwrap()
	return renderable.Renderable.text
}

func Layout(co layout.Constraint, text string) geometry.Size {
	return geometry.Size{Width: len(text), Height: 1}
}

func Draw(surface draw.Surface, text string, style draw.CellStyle) {
	// TODO: iterate over grapheme instead of runes.
	for i, c := range text {
		surface.Set(geometry.Vec2D{X: i, Y: 0}, draw.Cell{
			Style:   style,
			Content: c,
		})
	}
}
