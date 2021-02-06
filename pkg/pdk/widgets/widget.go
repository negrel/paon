package widgets

import (
	"fmt"
	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/internal/tree"
	"github.com/negrel/paon/pkg/pdk/events"
	"github.com/negrel/paon/pkg/pdk/styles"
	"github.com/negrel/paon/pkg/pdk/widgets/themes"
)

// Widget define any object part of the Widget tree
// that can be rendered in the screen.
type Widget interface {
	fmt.Stringer
	tree.Node
	events.Target
	styles.Stylised
	themes.Themed

	// ID return the unique generated ID or the given one using the ID Option.
	ID() string

	// Root return the root Widget in the tree.
	Root() Root

	// Parent return the Layout that contain this Widget in the tree.
	Parent() Layout

	// Next return the next sibling of the Widget.
	Next() Widget

	// Previous return the previous sibling of the Widget.
	Previous() Widget
}

var _ Widget = &widget{}

type widget struct {
	tree.Node
	events.Target

	theme themes.Theme
	id    string
}

// NewWidget return a new Widget object customized with the given Option.
func NewWidget(opts ...Option) Widget {
	w := newWidget(tree.NewNode())

	for _, opt := range opts {
		opt(w)
	}

	return w
}

func newWidget(node tree.Node) *widget {
	w := &widget{
		Node:   node,
		Target: events.MakeTarget(),
	}
	w.theme = themes.Make(func() themes.Themed { return w.Parent() })
	w.theme.AddEventListener(themes.ThemeChangeListener(func(_ themes.EventThemeChange) {

	}))

	return w
}

// ID implements the Widget interface.
func (w *widget) ID() string {
	return w.id
}

// String implements fmt.Stringer interface.
func (w *widget) String() string {
	return w.ID()
}

// Root implements the Widget interface.
func (w *widget) Root() Root {
	if r := w.RootNode(); r != nil {
		return r.(Root)
	}

	return nil
}

// Parent implements the Widget interface.
func (w *widget) Parent() Layout {
	if p := w.ParentNode(); p != nil {
		return p.(Layout)
	}
	return nil
}

// Next implements the Widget interface.
func (w *widget) Next() Widget {
	if n := w.NextNode(); n != nil {
		return n.(Widget)
	}

	return nil
}

// Previous implements the Widget interface.
func (w *widget) Previous() Widget {
	if p := w.PreviousNode(); p != nil {
		return p.(Widget)
	}

	return nil
}

// Style implements the styles.Stylised interface.
func (w *widget) Style() styles.Style {
	return w.theme
}

// Theme return the theme of the widget.
func (w *widget) Theme() themes.Theme {
	return w.theme
}

// ParentObject implements the render.Object interface.
func (w *widget) ParentObject() render.Object {
	return w.Parent()
}

func (w *widget) Size() geometry.Size {
	return geometry.NewSize(w.Width(), w.Height())
}

func (w *widget) cleanCache() {
	w.cache = cache{}
}

func (w *widget) Width() int {
	var width int
	if w.cache.validWidth {
		width = w.cache.width
	} else {
		width = w.width()
		w.cache.width = width
	}

	return width
}

func (w *widget) width() int {
	return w.computeHeightOrWidth(property.IDWidth, property.IDMinWidth, property.IDMaxWidth)
}

func (w *widget) Height() int {
	var height int

	if w.cache.validHeight {
		height = w.cache.height
	} else {
		height = w.height()
		w.cache.height = height
	}

	return height
}

func (w *widget) height() int {
	return w.computeHeightOrWidth(property.IDHeight, property.IDMinHeight, property.IDMaxHeight)
}

func (w *widget) computeHeightOrWidth(p, min, max property.ID) int {
	result := -1

	if r := w.theme.Get(p); w != nil {
		result = w.toCellUnitValue(r.(property.Unit).Unit, p == property.IDWidth)
	}

	if max := w.theme.Get(min); max != nil {
		maxR := max.(property.Unit).Value
		result = math.Min(result, maxR)
	}

	if min := w.theme.Get(max); min != nil {
		minR := min.(property.Unit).Value
		result = math.Max(result, minR)
	}

	return result
}

func (w *widget) toCellUnitValue(uv value.Unit, width bool) int {
	switch uv.ID {
	case value.CellUnit:
		return uv.Value

	case value.PercentageUnit:
		parent := w.Parent()
		if parent == nil {
			return -1
		}

		if width {
			return parent.Width() / 100 * uv.Value
		}
		return parent.Height() / 100 * uv.Value

	case value.WindowWidthUnit:
		return w.Root().Width() / 100 * uv.Value

	case value.WindowHeightUnit:
		return w.Root().Height() / 100 * uv.Value

	default:
		panic("can't convert unknown unit value to cell unit")
	}
}

func (w *widget) Position() geometry.Point {
	var position geometry.Point
	if w.cache.validPosition {
		position = w.cache.position
	} else {
		position = w.position()
		w.cache.position = position
	}

	return position
}

func (w *widget) position() geometry.Point {
	return geometry.Point{}
}
