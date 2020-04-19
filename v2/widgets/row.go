package widgets

import (
	"image"
	"log"

	"github.com/negrel/ginger/v2/render"
)

var _ Layout = &_row{}
var _ Widget = &_row{}

// _row is a layout that arrange widget horizontally.
type _row struct {
	*CoreLayout
}

// Row return a layout that arrange widget horizontally.
func Row(children []Widget) Layout {
	row := &_row{
		CoreLayout: NewCoreLayout("row", children),
	}

	row.Rendering = row.rendering

	return row
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Widget interface

func (r *_row) renderChilds(co Constraint) ([]*render.Frame, image.Point) {
	childCount := len(r.Children)
	childrenFrames := make([]*render.Frame, childCount)
	size := image.Pt(0, 0)

	for i := 0; i < childCount; i++ {
		childFrame := r.Children[i].Render(co)
		childrenFrames[i] = childFrame

		if cHeight := childFrame.Patch.Height(); cHeight > size.Y {
			size.Y = cHeight
		}

		// updating co.Bounds for next children
		co.Bounds.Min.X += childFrame.Patch.Width()
		// updating total width
		size.X += childFrame.Patch.Width()
	}

	return childrenFrames, size
}

// Rendering implements Widget interface.
func (r *_row) rendering(co Constraint) *render.Frame {
	// children constraint are relative
	childConstraint := Constraint{
		image.Rectangle{
			Min: image.Pt(0, 0),
			Max: co.Bounds.Max.Sub(co.Bounds.Min),
		},
	}
	childrenFrames, size := r.renderChilds(childConstraint)

	frame := render.NewFrame(co.Bounds.Min, size.X, size.Y)

	for i, childFrame := range childrenFrames {
		err := frame.Add(childFrame)

		if err != nil {
			log.Print("ROW:", err)
			log.Printf("ROW FRAME: %+v %+v %+v", frame.Position, frame.Patch.Width(), frame.Patch.Height())
			log.Fatalf("CHILD n°%v FRAME: %+v %+v %+v", i, childFrame, childFrame.Patch.Width(), childFrame.Patch.Height())
		}
	}

	return frame
}
