package widgets

import (
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/packer"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/layout"
	"github.com/negrel/paon/pdk/render"
	"github.com/negrel/paon/pdk/widgets"
	pdkwidgets "github.com/negrel/paon/pdk/widgets"
	"github.com/negrel/paon/props"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/styles/property"
	"github.com/negrel/paon/styles/value"
)

var _ pdkwidgets.Widget = &Text{}

var textDefaultStyle = styles.New()

func init() {
	textDefaultStyle.Set(
		props.FgColor(value.ColorFromHex(0xFFFFFF)),
	)
}

// Text define a basic text widget.
type Text struct {
	*pdkwidgets.BaseWidget

	content string
	cache   *layout.Cache
}

// NewText returns a basic text widget.
func NewText(content string) *Text {
	text := &Text{
		content: content,
	}
	text.cache = layout.NewCache(packer.NewText(&text.content))
	text.BaseWidget = pdkwidgets.NewBaseWidget(
		widgets.DefaultStyle(textDefaultStyle),
		widgets.Wrap(text),
		widgets.Renderable(text),
	)

	return text
}

// Content returns the content of this widget.
func (t Text) Content() string {
	return t.content
}

// SetContent sets the content of this Text widget.
func (t *Text) SetContent(content string) {
	t.content = content
	t.cache.Expire()
	pdkwidgets.NeedRender(t.BaseWidget)
}

// Layer implements the render.Renderable interface.
func (t Text) Layer() *render.Layer {
	if parent := t.Parent(); parent != nil {
		return parent.Layer()
	}

	return nil
}

// Render implements the render.Renderable interface.
func (t *Text) Render(ctx render.Context) {
	ctx.Flags = render.FullRenderFlag
	needRedraw := pdkwidgets.IsDirty(t.BaseWidget) || ctx.Flags >= render.DrawFlag
	needReflow := !t.cache.IsValid(ctx.Constraint) || ctx.Flags >= render.FullRenderFlag
	if !needRedraw && !needReflow {
		return
	}

	var box layout.BoxModel
	if needReflow {
		box = t.cache.Pack(ctx.Constraint)
	}
	if ctx.Flags >= render.FullRenderFlag {
		origin := ctx.Layout.Layout(box.MarginBox()).Min
		t.cache.SetPosition(origin)
	}
	box = t.cache.BoxModel()

	srf := ctx.Layer.SubSurface(box.MarginBox())
	color := t.Theme().Get(property.ForegroundColorID()).(property.Color)
	for i, char := range t.content {
		srf.Set(geometry.NewVec2D(i, 0), draw.Cell{
			Style: draw.CellStyle{
				Foreground: color.Color,
			},
			Content: char,
		})
	}
	pdkwidgets.MarkAsClean(t.BaseWidget)
}
