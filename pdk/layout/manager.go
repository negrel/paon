package layout

// Manager calculates a BoxModel based on the given constraint.
type Manager interface {
	Layout(Constraint) *Box
}

var _ Manager = ManagerFn(nil)

// ManagerFn is a simple function that implements the Manager interface.
type ManagerFn func(Constraint) *Box

// Layout implements the Manager interface.
func (mf ManagerFn) Layout(c Constraint) *Box {
	return mf(c)
}
