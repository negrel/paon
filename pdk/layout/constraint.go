package layout

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pdk/math"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/styles/property"
	"github.com/negrel/paon/styles/value"
)

// Constraint define the constraint an flow must respect.
type Constraint struct {
	MinSize, MaxSize geometry.Size
	ParentSize       geometry.Size
	RootSize         geometry.Size
}

// SetMin sets the constraint minimum rectangle.
func (c Constraint) SetMin(min geometry.Size) Constraint {
	c.MinSize = min
	return c
}

// SetMax sets the constraint maximum rectangle.
func (c Constraint) SetMax(max geometry.Size) Constraint {
	c.MaxSize = max
	return c
}

// Equals returns true if the given Constraint is equal to this Constraint.
func (c Constraint) Equals(other Constraint) bool {
	return c.MinSize.Equals(other.MinSize) && c.MaxSize.Equals(other.MaxSize) &&
		c.ParentSize.Equals(other.ParentSize) && c.RootSize.Equals(other.RootSize)
}

// ToCellUnit converts the given value to a Cell based value.Unit and returns it.
func (c Constraint) ToCellUnit(unit value.Unit) int {
	switch unit.UnitID {
	case value.CellUnit:
		return unit.Value

	case value.PercentageWidthUnit:
		return c.ParentSize.Width() / 100 * unit.Value

	case value.PercentageHeightUnit:
		return c.ParentSize.Height() / 100 * unit.Value

	case value.WindowWidthUnit:
		return c.RootSize.Width() / 100 * unit.Value

	case value.WindowHeightUnit:
		return c.RootSize.Height() / 100 * unit.Value

	default:
		panic("invalid unit")
	}
}

// NewFromStyle returns a stricter Constraint based on the given style.
// The new constraint will either use MinWidth/MinHeight property as minimum
// size or the current constraint size. The same is done for maximum size.
func (c Constraint) NewFromStyle(style styles.Style) Constraint {
	newc := Constraint{
		ParentSize: c.ParentSize,
		RootSize:   c.RootSize,
	}

	minWidth := c.MinSize.Width()
	maxWidth := c.MaxSize.Width()
	minWidthProp, ok := UnitProp(style, property.MinWidthID())
	if ok {
		minWidth = c.ApplyOnWidthProp(minWidthProp)
	}
	maxWidthProp, ok := UnitProp(style, property.MaxWidthID())
	if ok {
		maxWidth = c.ApplyOnWidthProp(maxWidthProp)
	}

	minHeight := c.MinSize.Height()
	maxHeight := c.MaxSize.Height()
	minHeightProp, ok := UnitProp(style, property.MinHeightID())
	if ok {
		minHeight = c.ApplyOnHeightProp(minHeightProp)
	}
	maxHeightProp, ok := UnitProp(style, property.MaxHeightID())
	if ok {
		maxHeight = c.ApplyOnHeightProp(maxHeightProp)
	}

	newc.MinSize = geometry.NewSize(minWidth, minHeight)
	newc.MaxSize = geometry.NewSize(maxWidth, maxHeight)

	return newc
}

// ApplyOnSize applies the size constraints on the given size and return
// the constrained size.
func (c Constraint) ApplyOnSize(size geometry.Size) geometry.Size {
	return geometry.NewSize(
		c.ApplyOnWidth(size.Width()),
		c.ApplyOnHeight(size.Height()),
	)
}

// ApplyOnWidthProp applies the width constraints on the given width property and
// return the constrained width as an int (CellUnit).
func (c Constraint) ApplyOnWidthProp(widthProp property.Unit) int {
	width := c.ToCellUnit(widthProp.Value)
	return c.ApplyOnWidth(width)
}

// ApplyOnWidth applies the width constraints on the given width and return the
// constrained width.
func (c Constraint) ApplyOnWidth(width int) int {
	width = math.Max(width, c.MinSize.Width())
	return math.Min(width, c.MaxSize.Width())
}

// ApplyOnHeightProp applies the height constraints on the given height property and
// return the constrained height as an int (CellUnit).
func (c Constraint) ApplyOnHeightProp(heightProp property.Unit) int {
	height := c.ToCellUnit(heightProp.Value)
	return c.ApplyOnHeight(height)
}

// ApplyOnHeight applies the height constraints on the given height and return the
// constrained height.
func (c Constraint) ApplyOnHeight(height int) int {
	height = math.Max(height, c.MinSize.Height())
	return math.Min(height, c.MaxSize.Height())
}
