package composite

import (
	"bytes"
	"io"
	"sync"

	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/styles/value"
)

func makeBuf(rect geometry.Rectangle) []draw.Cell {
	return make([]draw.Cell, rect.Area())
}

var _ draw.Canvas = Layer{}

// Layer is an implementation of the draw.Canvas interface.
// Layers are then merged before the painting to the backend surface.
type Layer struct {
	mu *sync.Mutex

	cells  []draw.Cell
	bounds geometry.Rectangle
}

// NewLayer returns a new Layer with the given bounds.
func NewLayer(bounds geometry.Rectangle) Layer {
	return Layer{
		mu:     &sync.Mutex{},
		cells:  makeBuf(bounds),
		bounds: bounds,
	}
}

// Bounds implements the draw.Canvas interface.
func (l Layer) Bounds() geometry.Rectangle {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.bounds
}

func (l Layer) index(pt geometry.Point) int {
	return pt.Y()*l.bounds.Width() + pt.X()
}

// Get implements the draw.Canvas interface.
func (l Layer) Get(pt geometry.Point) draw.Cell {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.get(pt)
}

func (l Layer) get(pt geometry.Point) draw.Cell {
	pt = pt.Add(l.bounds.Min)
	if l.bounds.Contains(pt) {
		return l.cells[l.index(pt)]
	}

	return draw.ZeroCell()
}

// Set implements the draw.Canvas interface.
func (l Layer) Set(pt geometry.Point, cell draw.Cell) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.set(pt, cell)
}

func (l Layer) set(pt geometry.Point, cell draw.Cell) {
	pt = pt.Add(l.bounds.Min)

	if l.bounds.Contains(pt) {
		l.cells[l.index(pt)] = cell
	}
}

// Rebounds reposition and resize the layer.
func (l Layer) Rebounds(new geometry.Rectangle) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// We don't copy old cells since everything must be repaint
	l.cells = makeBuf(new)
	l.bounds = new
}

// Dump the layer to the given io writer
func (l Layer) Dump(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()

	var buf bytes.Buffer

	for i, cell := range l.cells {
		if i%l.bounds.Width() == 0 {
			buf.WriteRune('\n')
			w.Write(buf.Bytes())
			buf.Reset()
		}
		buf.WriteRune(cell.Content)
	}
}

// Composite merge the given layers into the given draw.Canvas.
// Layers are merged into the given orders, therefore you must sort them
// if you have overlapping region.
func Composite(dst draw.Canvas, layers ...Layer) {
	// TODO redraw only damaged areas
	for _, layer := range layers {
		bounds := layer.Bounds()

		for x := bounds.Min.X(); x < bounds.Max.X(); x++ {
			for y := bounds.Min.Y(); y < bounds.Max.Y(); y++ {

			}
		}
	}
}

func compose(a, b draw.Cell) draw.Cell {
	content := b.Content
	style := b.Style
	if b.Content == '\000' {
		style = a.Style
		content = a.Content
	}

	return draw.Cell{
		Style:   style,
		Content: content,
	}
}

func composeColor(a, b value.Color) value.Color {
	return b
}
