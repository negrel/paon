package widget

import (
	"image"
	"log"

	"github.com/negrel/ginger/v1/style"
)

var _ Widget = &Row{}

// Row is a layout that arrange widget horizontaly.
type Row struct {
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
func (r *Row) drawChild(c Constraint) (childsFrames []*style.Frame, totalW, maxH int) {
	childCount := len(r.Childrens)
	childsFrames = make([]*style.Frame, childCount)

	// Computing child frame
	for i := 0; i < childCount; i++ {
		child := r.Childrens[i]

		// Child frame
		cFrame := child.Draw(c)
		cWidth := cFrame.R.Dx() // Child Width
		childsFrames[i] = cFrame

		// Reducing constraint freespace
		c.R.Min = c.R.Min.Add(image.Point{cWidth, 0})

		// Updating total & max
		totalW += cWidth
		if cHeight := cFrame.R.Dy(); cHeight > maxH {
			maxH = cHeight
		}
	}

	return
}

// Draw implements the widget interface
func (r *Row) Draw(c Constraint) *style.Frame {
	childsFrames, _, rHeight := r.drawChild(c)

	frame := style.NewFrame(0, rHeight)

	log.Printf("%+v", frame)

	// Format frame for the same width
	for i := 0; i < len(childsFrames); i++ {
		childsFrames[i].SetHeight(rHeight, c.c)
		frame.AppendRight(childsFrames[i].G)
	}

	return frame
}
