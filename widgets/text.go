package widgets

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/layout"
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
}

// NewText returns a basic text widget.
func NewText(content string) *Text {
	text := &Text{
		content: content,
	}
	text.BaseWidget = pdkwidgets.NewBaseWidget(
		widgets.LayoutManager(layout.ManagerFn(func(c layout.Constraint) *layout.Box {
			return layout.NewBox(geometry.NewSize(len(text.content), 1))
		})),
		widgets.Drawer(TextDrawer(text, text.Content)),
		widgets.DefaultStyle(textDefaultStyle),
		widgets.Wrap(text),
	)

	return text
}

// Content returns the content of this widget.
func (t *Text) Content() string {
	return t.content
}

// SetContent sets the content of this Text widget.
func (t *Text) SetContent(content string) {
	t.content = content
	pdkwidgets.Reflow(t.BaseWidget, true)
}

// TextDrawer returns a drawer that draw the text returned by getContent
// with the styles of the given styled.
func TextDrawer(styled styles.Styled, getContent func() string) draw.Drawer {
	return draw.DrawerFn(func(c draw.Canvas) {
		style := styled.Style()

		color := style.Get(property.ForegroundColorID()).(property.Color)
		draw.NewContext(c).
			FillTextH(geometry.NewVec2D(0, 0), draw.CellStyle{Foreground: color.Color}, getContent())
	})
}
