package widgets

import (
	"github.com/negrel/debuggo/code_gen/log/log"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/packer"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/pdk/layout"
	"github.com/negrel/paon/pdk/render"
	"github.com/negrel/paon/pdk/widgets"
	pdkwidgets "github.com/negrel/paon/pdk/widgets"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/styles/property"
)

var _ pdkwidgets.Widget = &Text{}

var textDefaultStyle = styles.New(events.NewTarget())

func init() {
	textDefaultStyle.SetColor(
		property.ForegroundColor(),
		&property.White,
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
	needRedraw := pdkwidgets.IsDirty(t.BaseWidget) || ctx.Flags|render.DrawFlag != 0
	needResize := !t.cache.IsValid(ctx.Constraint)
	needReflow := ctx.Flags|render.LayoutFlag != 0
	if !needRedraw && !needReflow && !needResize {
		return
	}

	var box layout.BoxModel
	if needResize {
		box = t.cache.Pack(ctx.Constraint)
	} else {
		box = t.cache.BoxModel()
	}

	if needReflow {
		origin := ctx.Layout.Layout(box.MarginBox()).Min
		t.cache.SetPosition(origin)
	}
	box = t.cache.BoxModel() // Positionned box

	var srf draw.SubSurface
	if prop := t.Theme().Int(property.ZIndex()); prop != nil {
		layer := render.NewLayer(t.Node())
		err := ctx.Layer.AddLayer(layer)
		if err != nil {
			log.Error(err)
		}
		srf = layer.SubSurface(box.MarginBox())
	} else {
		srf = ctx.Layer.SubSurface(box.MarginBox())
	}

	var c property.Color
	if fgColor := t.Theme().Color(property.ForegroundColor()); fgColor != nil {
		c = *fgColor
	}

	// bgColor := t.Theme().Color(property.BackgroundColor())
	for i, char := range t.content {
		srf.Set(geometry.NewVec2D(i, 0), draw.Cell{
			Style: draw.CellStyle{
				Foreground: c,
				// Background: *bgColor,
			},
			Content: char,
		})
	}
	pdkwidgets.MarkAsClean(t.BaseWidget)
}
