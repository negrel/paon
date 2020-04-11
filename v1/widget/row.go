package widget

import (
	"image"
	"log"

	"github.com/negrel/ginger/v1/painting"
)

var _ Layout = &_row{}
var _ Widget = &_row{}

// _row is a layout that arrange widget horizontally.
type _row struct {
	*LayoutMultipleChild
}

// Row return a layout that arrange widget horizontally.
func Row(Children []Widget) Layout {
	row := &_row{
		&LayoutMultipleChild{
			Core:     &Core{},
			Children: Children,
		},
	}

	row.ForEach(func(_ int, child Widget) {
		child.AdoptedBy(row)
	})

	return row
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Widget interface

func (r *_row) drawChilds(bounds image.Rectangle) ([]*painting.Frame, image.Point) {
	childCount := len(r.Children)
	cFrames := make([]*painting.Frame, childCount)
	size := image.Pt(0, 0)

	for i := 0; i < childCount; i++ {
		cFrame := r.Children[i].Draw(bounds)
		cFrames[i] = cFrame

		if cHeight := cFrame.Patch.Height(); cHeight > size.Y {
			size.Y = cHeight
		}

		// updating bounds for next children
		bounds.Min.X += cFrame.Patch.Width()
		// updating total width
		size.X += cFrame.Patch.Width()
	}

	return cFrames, size
}

// Draw implements Widget interface.
func (r *_row) Draw(bounds image.Rectangle) *painting.Frame {
	// child bounds are relative
	cBounds := image.Rectangle{
		Min: image.Point{},
		Max: bounds.Max.Sub(bounds.Min),
	}
	cFrames, size := r.drawChilds(cBounds)

	frame := painting.NewFrame(bounds.Min, size.X, size.Y)

	for i := 0; i < len(cFrames); i++ {
		err := frame.Add(cFrames[i])

		if err != nil {
			log.Print("ROW:", err)
			log.Printf("ROW FRAME: %+v %+v %+v", frame.Position, frame.Patch.Width(), frame.Patch.Height())
			log.Fatalf("CHILD n°%v FRAME: %+v %+v %+v", i, cFrames[i], cFrames[i].Patch.Width(), cFrames[i].Patch.Height())
		}
	}

	return frame
}
