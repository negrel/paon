package node

var _ Leaf = &BaseLeaf{}

// BaseLeaf is intended to be integrated into the
// elements of the node tree.
type BaseLeaf struct {
	parent Branch

	owner Branch
}

/*****************************************************
 ***************** Getters & Setters *****************
 *****************************************************/
// ANCHOR Getters & Setters

// Attached implements Leaf interface.
func (l *BaseLeaf) Attached() bool {
	return l.owner != nil
}

// Attach implements Leaf interface.
func (l *BaseLeaf) Attach(owner Branch) {
	l.owner = owner
}

// Detach implements Leaf interface.
func (l *BaseLeaf) Detach() {
	l.owner = nil
}

// Owner implements Leaf interface.
func (l *BaseLeaf) Owner() Branch {
	return l.owner
}

// Parent implements Leaf interface.
func (l *BaseLeaf) Parent() Branch {
	return l.parent
}

func (l *BaseLeaf) setParent(parent Branch) {
	l.parent = parent
}
