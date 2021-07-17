package layout

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/styles"
)

// Boxed define any objects that can returns a BoxModel.
type Boxed interface {
	Box() BoxModel
}

// BoxModel define a sized box with margin, border and padding in a 2D geometric plane.
type BoxModel interface {
	MarginBox() geometry.Rectangle
	BorderBox() geometry.Rectangle
	PaddingBox() geometry.Rectangle
	ContentBox() geometry.Rectangle
}

var _ BoxModel = movedBox{}

// movedBox is a wrapper around a BoxModel object.
// This allow layout optimization for object that has the same size constraint
// but has moved.
type movedBox struct {
	BoxModel
	offset geometry.Point
}

func (mb movedBox) MarginBox() geometry.Rectangle {
	return mb.BoxModel.MarginBox().Add(mb.offset)
}

func (mb movedBox) BorderBox() geometry.Rectangle {
	return mb.BoxModel.BorderBox().Add(mb.offset)
}

func (mb movedBox) PaddingBox() geometry.Rectangle {
	return mb.BoxModel.PaddingBox().Add(mb.offset)
}

func (mb movedBox) ContentBox() geometry.Rectangle {
	return mb.BoxModel.ContentBox().Add(mb.offset)
}

// Box define a basic BoxModel implementation.
type Box struct {
	borderBox geometry.Rectangle
	marginBoxOffset,
	paddingBoxOffset,
	contentBoxOffset boxOffset
}

// NewBox return a new Box with the given content box.
func NewBox(rectangle geometry.Rectangle) *Box {
	return &Box{
		borderBox: rectangle,
	}
}

// MarginBox implements the BoxModel interface.
func (b *Box) MarginBox() geometry.Rectangle {
	return b.marginBoxOffset.applyOn(b.borderBox)
}

// BorderBox implements the BoxModel interface.
func (b *Box) BorderBox() geometry.Rectangle {
	return b.borderBox
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
	marginBox := b.MarginBox()
	diffW := size.Width() - marginBox.Width()
	diffH := size.Height() - marginBox.Height()

	b.borderBox.Max = b.borderBox.Max.Add(
		geometry.Pt(diffW, diffH),
	)
}

// ApplyMargin applies the margin of the given style to the box.
func (b *Box) ApplyMargin(style styles.Style) {
	b.applyMargin(marginOf(style))
}

func (b *Box) applyMargin(margin boxOffset) {
	b.borderBox.Min = b.borderBox.Min.Add(
		geometry.Pt(b.marginBoxOffset.left(), b.marginBoxOffset.top()),
	)
	b.borderBox.Max = b.borderBox.Max.Add(
		geometry.Pt(b.marginBoxOffset.left(), b.marginBoxOffset.top()),
	)

	b.marginBoxOffset = margin.reverse()

	b.borderBox.Min = b.borderBox.Min.Add(
		geometry.Pt(margin.left(), margin.top()),
	)
	b.borderBox.Max = b.borderBox.Max.Add(
		geometry.Pt(margin.left(), margin.top()),
	)
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
