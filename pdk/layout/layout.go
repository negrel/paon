package layout

// Layout is a generic interface for object
// that can produce BoxModel using the given Constraint.
type Layout interface {
	Layout(Constraint) BoxModel
}
