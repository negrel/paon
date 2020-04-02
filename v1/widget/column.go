package widget

import (
	"image"

	"github.com/negrel/ginger/v1/style"
)

var _ Widget = &Column{}

// Column is a layout that arrange widget vertically.
type Column struct {
	*Base

	Childrens []Widget
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Widget interface

// drawChild return the childrens frame with the total width
// and the max height.
func (col *Column) drawChild(c Constraint) (childsFrames []*style.Frame, totalH, maxW int) {
	childCount := len(col.Childrens)
	childsFrames = make([]*style.Frame, childCount, childCount)
	// Computing child frame
	for i := 0; i < childCount; i++ {
		child := col.Childrens[i]

		// Child frame
		cFrame := child.Draw(c)
		cHeight := cFrame.R.Dy() // Child height
		childsFrames[i] = cFrame

		// Reducing constraint freespace for next child
		c.R.Min = c.R.Min.Add(image.Pt(0, cHeight))

		// Updating total & max
		totalH += cHeight
		if cWidth := cFrame.R.Dx(); cWidth > maxW {
			maxW = cWidth
		}
	}

	return
}

// Draw implements the widget interface
func (col *Column) Draw(c Constraint) *style.Frame {
	childsFrames, _, colWidth := col.drawChild(c)

	frame := style.NewFrame(colWidth, 0)

	// Format frame for the same width
	for i := 0; i < len(childsFrames); i++ {
		childsFrames[i].SetWidth(colWidth, c.c)
		frame.AppendBelow(childsFrames[i].G)
	}

	return frame
}
