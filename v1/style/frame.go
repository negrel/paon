package style

import (
	"image"
)

// Frame represent raw data to paint
type Frame struct {
	R image.Rectangle
	G Grid
}

// NewFrame return a new frame
func NewFrame(width, height int) *Frame {
	var g Grid
	for i := 0; i < height; i++ {
		var r = make(Row, 0, width)

		for j := 0; j < width; j++ {
			r = append(r, &Cell{})
		}
		g = append(g, r)
	}

	return &Frame{
		R: image.Rectangle{
			Min: image.Pt(0, 0),
			Max: image.Pt(width, height),
		},
		G: g,
	}
}

/*****************************************************
 ***************** Getters & Setters *****************
 *****************************************************/
// ANCHOR Getters & Setters

// SetWidth method make grow the frame width to the given
// width. Then it grow the cells grid using empty cells of
// the given colors.
func (f *Frame) SetWidth(w int, c Colors) {
	// Diff width
	wDiff := w - f.R.Dx()
	// Update rectange frame
	f.R.Max.X += wDiff

	// The empty cell use to fill
	empty := &Cell{
		Char:   ' ',
		Colors: c,
	}

	// Append column to match the width
	col := NewRow(wDiff, empty)
	for i := 0; i < f.R.Dy(); i++ {
		f.AppendAt(col, i)
	}
}

// SetHeight method make grow the frame height to the given
// height. Then it grow the cells grid using empty cells of
// the given colors.
func (f *Frame) SetHeight(h int, c Colors) {
	// Diff height
	hDiff := h - f.R.Dy()
	// Update rectange frame
	f.R.Max.Y += hDiff

	// The empty cell use to fill
	empty := &Cell{
		Char:   ' ',
		Colors: c,
	}

	// Append column to match the height
	row := NewRow(hDiff, empty)
	for i := 0; i < hDiff; i++ {
		f.G.appendRow(row)
	}
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// AppendAt method append the given row to the n grid
// row.
func (f *Frame) AppendAt(r Row, n int) {
	f.R.Max = f.R.Max.Add(image.Pt(len(r), 0))
	f.G.appendAt(r, n)
}

// AppendRow method append the given row to the
// frame grid and update the rectangle
func (f *Frame) AppendRow(r Row) {
	f.R.Max = f.R.Max.Add(image.Pt(0, 1))
	f.G.appendRow(r)
}

// AppendBelow method append the given grid below
func (f *Frame) AppendBelow(g Grid) {
	for i := 0; i < len(g); i++ {
		f.AppendRow(g[i])
	}
}

// AppendRight method append the given grid at the right.
func (f *Frame) AppendRight(g Grid) {
	for i := 0; i < len(g); i++ {
		f.AppendAt(g[i], i)
	}
}
