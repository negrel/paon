package layout

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/styles"
)

// Object define objects that can return a size boxed
// based on the given constraint. The returned box can
// then be positionned by a layout.
type Object interface {
	Layout(Constraint) BoxModel
}

// ObjectFunc define a function that implements the Object interface.
type ObjectFunc func(Constraint) BoxModel

// Layout implements the Object interface.
func (of ObjectFunc) Layout(c Constraint) BoxModel {
	return of(c)
}

// BoxedObject define an object with a BoxModel.
type BoxedObject interface {
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
	boxSize geometry.Size
	borderBoxOffset,
	paddingBoxOffset,
	contentBoxOffset boxOffset
}

// NewBox return a new Box with the given content box.
func NewBox(size geometry.Size) Box {
	return Box{
		boxSize: size,
	}
}

// MarginBox implements the BoxModel interface.
func (b Box) MarginBox() geometry.Rectangle {
	return geometry.Rectangle{
		Min: geometry.Vec2D{},
		Max: geometry.NewVec2D(b.boxSize.Width(), b.boxSize.Height()),
	}
}

// BorderBox implements the BoxModel interface.
func (b Box) BorderBox() geometry.Rectangle {
	return b.borderBoxOffset.applyOn(b.MarginBox())
}

// PaddingBox implements the BoxModel interface.
func (b Box) PaddingBox() geometry.Rectangle {
	return b.paddingBoxOffset.applyOn(b.BorderBox())
}

// ContentBox implements the BoxModel interface.
func (b Box) ContentBox() geometry.Rectangle {
	return b.contentBoxOffset.applyOn(b.PaddingBox())
}

// Resize change the size of the margin box.
func (b Box) Resize(size geometry.Size) {
	b.boxSize = size
}

// ApplyMargin applies the margin of the given style to the box.
func (b Box) ApplyMargin(style styles.Style) Box {
	b.applyMargin(marginOf(style))
	return b
}

func (b Box) applyMargin(margin boxOffset) {
	b.borderBoxOffset = margin
}

// ApplyBorder applies the border of the given style to the box.
func (b Box) ApplyBorder(style styles.Style) Box {
	b.applyBorder(borderOf(style))
	return b
}

func (b Box) applyBorder(border boxOffset) {
	b.paddingBoxOffset = border
}

// ApplyPadding applies the padding of the given style to the box.
func (b Box) ApplyPadding(style styles.Style) Box {
	b.applyPadding(paddingOf(style))
	return b
}

func (b Box) applyPadding(padding boxOffset) {
	b.contentBoxOffset = padding
}
