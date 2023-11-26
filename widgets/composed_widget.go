package widgets

import (
	"github.com/negrel/paon/render"
	"github.com/negrel/paon/styles"
)

// ComposedWidget define a Widget composed of a PanicWidget,
// a styles.Styled and a render.RenderableAccessor.
type ComposedWidget struct {
	*PanicWidget
	styles.Styled
	render.RenderableAccessor
}

func NewComposedWidget(styled styles.Styled, renderableAccessor render.RenderableAccessor) *ComposedWidget {
	w := &ComposedWidget{
		Styled:             styled,
		RenderableAccessor: renderableAccessor,
	}
	w.PanicWidget = NewPanicWidget(w)

	return w
}

// Style implements styles.Styled.
func (cw ComposedWidget) Style() styles.Style {
	return cw.Styled.Style()
}

// Renderable implements render.RenderableAccessor.
func (cw ComposedWidget) Renderable() render.Renderable {
	return cw.RenderableAccessor.Renderable()
}
