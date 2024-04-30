package layout

// Layout is a generic interface for sized objects.
type Layout[T any] interface {
	Layout(Constraint, T) SizeHint
}
