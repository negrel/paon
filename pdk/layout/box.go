package layout

import (
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/styles"
)

// BoxedObject define an object with a BoxModel.
type BoxedObject interface {
	BoxModel() BoxModel
}

// BoxModel define a box with margin, border and padding in a 2D geometric plane.
type BoxModel interface {
	MarginBox() geometry.Rectangle
	BorderBox() geometry.Rectangle
	PaddingBox() geometry.Rectangle
	ContentBox() geometry.Rectangle
}

var _ BoxModel = PositionedBoxModel{}
var _ geometry.Positioned = PositionedBoxModel{}

// PositionedBoxModel is a BoxModel wrapper that implement the
// geometry.Positioned interface.
type PositionedBoxModel struct {
	Origin geometry.Vec2D
	BoxModel
}

// Position implements the geometry.Positioned interface.
func (pbm PositionedBoxModel) Position() geometry.Vec2D {
	return pbm.Origin
}

// MarginBox implements the BoxModel interface.
func (pbm PositionedBoxModel) MarginBox() geometry.Rectangle {
	return pbm.BoxModel.MarginBox().MoveBy(pbm.Origin)
}

// BorderBox implements the BoxModel interface.
func (pbm PositionedBoxModel) BorderBox() geometry.Rectangle {
	return pbm.BoxModel.BorderBox().MoveBy(pbm.Origin)
}

// PaddingBox implements the BoxModel interface.
func (pbm PositionedBoxModel) PaddingBox() geometry.Rectangle {
	return pbm.BoxModel.PaddingBox().MoveBy(pbm.Origin)
}

// ContentBox implements the BoxModel interface.
func (pbm PositionedBoxModel) ContentBox() geometry.Rectangle {
	return pbm.BoxModel.ContentBox().MoveBy(pbm.Origin)
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
func (b *Box) MarginBox() geometry.Rectangle {
	return geometry.Rectangle{
		Min: geometry.Vec2D{},
		Max: geometry.NewVec2D(b.boxSize.Width(), b.boxSize.Height()),
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

// ApplyMargin applies the margin of the given style to the box.
func (b *Box) ApplyMargin(style styles.Style) {
	b.applyMargin(marginOf(style))
}

func (b *Box) applyMargin(margin boxOffset) {
	b.borderBoxOffset = margin
}

// ApplyBorder applies the border of the given style to the box.
func (b *Box) ApplyBorder(style styles.Style) {
	b.applyBorder(borderOf(style))
}

func (b *Box) applyBorder(border boxOffset) {
	b.paddingBoxOffset = border
}

// ApplyPadding applies the padding of the given style to the box.
func (b *Box) ApplyPadding(style styles.Style) {
	b.applyPadding(paddingOf(style))
}

func (b *Box) applyPadding(padding boxOffset) {
	b.contentBoxOffset = padding
}
