package draw

import (
	"github.com/negrel/paon/pdk/layout"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/styles/property"
)

// Drawer define an object that can draw on a Canvas.
type Drawer interface {
	Draw(Surface)
}

// DrawerFn define a function that implements the Drawer interface.
type DrawerFn func(Surface)

// Draw implements the Drawable interface.
func (fn DrawerFn) Draw(c Surface) {
	fn(c)
}

// Box returns a DrawerFn that draw the box of given boxed object with the given style.
// The contentDrawer is used to draw in the ContentBox.
func Box(styled styles.Styled, boxed layout.Boxed, contentDrawer Drawer) DrawerFn {
	return func(srf Surface) {
		style := styled.Style()
		box := boxed.Box()

		// Background color
		bgColor, ok := style.Get(property.BackgroundColorID()).(property.Color)
		if ok {
			ctx := NewContext(srf)
			ctx.FillRectangle(box.PaddingBox(), CellStyle{
				Background: bgColor.Color,
			})

			contentDrawer.Draw(NewSubSurface(srf, box.ContentBox()))
		}
	}
}
