package widgets

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/styles"
)

// ComposedWidget is Widget type made of composition. Unlike
// BaseWidget, you shouldn't overwrite drawing, layout and style methods.
// ComposedWidget contains no logic itself, it only forwards call to inner
// layout, drawer, or styled object.
type ComposedWidget struct {
	*BaseWidget
	Styled       styles.Styled
	LayoutLayout layout.Layout
	Drawer       draw.Drawer
}

func NewComposedWidget(
	styled styles.Styled,
	layout layout.Layout,
	drawer draw.Drawer,
) *ComposedWidget {
	fw := &ComposedWidget{
		Styled:       styled,
		LayoutLayout: layout,
		Drawer:       drawer,
	}

	fw.BaseWidget = NewBaseWidget(fw)

	return fw
}

// Style implements styles.Styled.
func (fw ComposedWidget) Style() styles.Style {
	return fw.Styled.Style()
}

// Layout implements layout.Layout.
func (fw ComposedWidget) Layout(co layout.Constraint) geometry.Size {
	return fw.LayoutLayout.Layout(co)
}

// Draw implements draw.Drawer.
func (fw ComposedWidget) Draw(surface draw.Surface) {
	fw.Drawer.Draw(surface)
}
