package layout

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/styles"
)

// Boxed define any objects that can returns a BoxModel.
type Boxed interface {
	Box() BoxModel
}

// BoxModel define a box with margin, border and padding in a 2D geometric plane.
type BoxModel interface {
	MarginBox() geometry.Rectangle
	BorderBox() geometry.Rectangle
	PaddingBox() geometry.Rectangle
	ContentBox() geometry.Rectangle
}

var _ BoxModel = &Box{}

// Box define a basic BoxModel implementation.
type Box struct {
	origin geometry.Point

	boxSize geometry.Size
	borderBoxOffset,
	paddingBoxOffset,
	contentBoxOffset boxOffset
}

// NewBox return a new Box with the given content box.
func NewBox(size geometry.Size) *Box {
	return &Box{
		boxSize: size,
	}
}

// MarginBox implements the BoxModel interface.
func (b *Box) MarginBox() geometry.Rectangle {
	return geometry.Rectangle{
		Min: b.origin,
		Max: b.origin.Add(geometry.Pt(b.boxSize.Width(), b.boxSize.Height())),
	}
}

// BorderBox implements the BoxModel interface.
func (b *Box) BorderBox() geometry.Rectangle {
	return b.borderBoxOffset.applyOn(b.MarginBox())
}

// PaddingBox implements the BoxModel interface.
func (b *Box) PaddingBox() geometry.Rectangle {
	return b.paddingBoxOffset.applyOn(b.BorderBox())
}

// ContentBox implements the BoxModel interface.
func (b *Box) ContentBox() geometry.Rectangle {
	return b.contentBoxOffset.applyOn(b.PaddingBox())
}

// Resize change the size of the margin box.
func (b *Box) Resize(size geometry.Size) {
	b.boxSize = size
}

// SetOrigin sets the origin coordinate of the top-left corner of the margin box.
func (b *Box) SetOrigin(pt geometry.Point) {
	b.origin = pt
}

// ApplyMargin applies the margin of the given style to the box.
func (b *Box) ApplyMargin(style styles.Style) *Box {
	b.applyMargin(marginOf(style))
	return b
}

func (b *Box) applyMargin(margin boxOffset) {
	b.borderBoxOffset = margin
}

// ApplyBorder applies the border of the given style to the box.
func (b *Box) ApplyBorder(style styles.Style) *Box {
	b.applyBorder(borderOf(style))
	return b
}

func (b *Box) applyBorder(border boxOffset) {
	b.paddingBoxOffset = border
}

// ApplyPadding applies the padding of the given style to the box.
func (b *Box) ApplyPadding(style styles.Style) *Box {
	b.applyPadding(paddingOf(style))
	return b
}

func (b *Box) applyPadding(padding boxOffset) {
	b.contentBoxOffset = padding
}
