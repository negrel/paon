package drawer

import (
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/layout"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/styles/property"
)

// NoDrawer is a draw.DrawerFn that don't draw anything.
func NoDrawer(ctx *draw.Context) {

}

// Box draws a given BoxModel with the given styles.Style on the given draw.Context.
func Box(styled styles.Styled, boxed layout.Boxed) draw.DrawerFn {
	return func(ctx *draw.Context) {
		style := styled.Style()
		box := boxed.Box()

		// Background color
		bgColor, ok := style.Get(property.BackgroundColorID()).(property.Color)
		if ok {
			ctx.SetFillColor(bgColor.Color)
		}
		ctx.FillRectangle(box.PaddingBox())
	}
}
