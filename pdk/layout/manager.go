package layout

// Manager calculates a BoxModel based on the given constraint.
type Manager interface {
	Layout(Constraint) BoxModel
}

var _ Manager = ManagerFn(nil)

// ManagerFn is a simple function that implements the Manager interface.
type ManagerFn func(Constraint) BoxModel

// Layout implements the Manager interface.
func (mf ManagerFn) Layout(c Constraint) BoxModel {
	return mf(c)
}
