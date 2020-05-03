package widgets

// Layout define any widgets that can containt child
// widgets.
type Layout interface {
	SingleChildLayout
	MultipleChildLayout
}
