package drawer

import (
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/layout"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/styles/property"
)

// Box draws a given BoxModel with the given styles.Style on the given draw.Context.
func Box(styled styles.Styled, boxed layout.Boxed) draw.DrawerFn {
	return func(c draw.Canvas) {
		style := styled.Style()
		box := boxed.Box()

		// Background color
		bgColor, ok := style.Get(property.BackgroundColorID()).(property.Color)
		if ok {
			ctx := draw.NewContext(c)
			ctx.FillRectangle(box.PaddingBox(), draw.CellStyle{
				Background: bgColor.Color,
			})
		}
	}

}
