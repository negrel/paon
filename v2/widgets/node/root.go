package node

var _ Branch = &Root{}

// Root is the root of the node tree. It should never have
// more than one root in an entire node tree.
type Root struct {
	*BaseBranch
}

// Attached implements Leaf interface.
func (r *Root) Attached() bool {
	return true
}
