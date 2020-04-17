package render

import (
	"errors"
	"image"
	"log"
)

// Frame are rectangular update screen patch
type Frame struct {

	// Position is the relative position where the
	// frame painting must start (from top left corner).
	Position Position

	// Patch is screen patch relative to the frame
	// position on the screen.
	Patch *Matrix
}

// NewFrame return a new Frame object
func NewFrame(p Position, width, height int) *Frame {
	return &Frame{
		Position: p,
		Patch:    NewMatrix(width, height),
	}
}

/*****************************************************
 ********************* Interface  ********************
 *****************************************************/
// ANCHOR Interface

// Paint implements Paintable interface.
func (f *Frame) Paint() [][]*RawCell {
	height := f.Patch.Height()
	width := f.Patch.Width()
	final := make([][]*RawCell, height)

	for i := 0; i < height; i++ {
		row := make([]*RawCell, width)

		yOffset := f.Position.Y + i

		for j := 0; j < width; j++ {
			xOffset := f.Position.X + j
			offset := image.Pt(xOffset, yOffset)

			row[j] = f.Patch.M[i][j].Compute(offset)
		}

		final[i] = row
	}

	return final
}

/*****************************************************
 ***************** GETTERS & SETTERS *****************
 *****************************************************/
// ANCHOR Getters & setter

// Bounds getter return the bounds of the frame.
func (f *Frame) Bounds() *image.Rectangle {
	return &image.Rectangle{
		Min: image.Point{
			X: f.Position.X,
			Y: f.Position.Y,
		},
		Max: image.Point{
			X: (f.Position.X + f.Patch.Width()),
			Y: (f.Position.Y + f.Patch.Height()),
		},
	}
}

// MaxHeightIndex return the max height index used by frame.
// (height + y position)
func (f *Frame) maxHeightIndex() uint {
	return uint(f.Position.Y + f.Patch.Height() - 1)
}

// MaxWidthIndex return the max width index used by frame.
// (Width + x position)
func (f *Frame) maxWidthIndex() uint {
	return uint(f.Position.X + f.Patch.Width() - 1)
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Add method add the given frame patch into the frame itself
// using the patch relative position.
func (f *Frame) Add(o *Frame) error {
	if can := f.CanContain(o); !can {
		return errors.New("the given frame can't be added or contained")
	}

	for i := 0; i < o.Patch.Height(); i++ {
		yOffset := o.Position.Y + i

		for j := 0; j < o.Patch.Width(); j++ {
			xOffset := o.Position.X + j
			f.Patch.M[yOffset][xOffset] = o.Patch.M[i][j]
		}
	}

	return nil
}

// CanContain return wether or not it can contain
// the given frame.
func (f *Frame) CanContain(o *Frame) bool {

	if !o.Patch.isValid() {
		log.Println("Can't contain because the frame is invalid.")
		return false
	}

	if o.maxHeightIndex() >= uint(f.Patch.Height()) {
		log.Println(o.maxHeightIndex() >= uint(f.Patch.Height()))

		log.Println("Can't contain because of the height.")
		log.Printf("Container height: %v", f.Patch.Height())
		log.Printf("Frame height: %v", o.maxHeightIndex())
		return false
	}

	if o.maxWidthIndex() >= uint(f.Patch.Width()) {
		log.Println(o.maxWidthIndex() >= uint(f.Patch.Width()))

		log.Println("Can't contain because of the width.")
		log.Printf("Container width:  %v", f.Patch.Width())
		log.Printf("Frame width: %v", o.maxWidthIndex())
		return false
	}

	return true
}

func (f *Frame) isEqual(other *Frame) bool {
	if !f.Patch.isEqual(other.Patch) {
		return false
	}

	return true
}
