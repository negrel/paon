package node

import "log"

var _ Branch = &BaseBranch{}

// BaseBranch is intended to be integrated into the
// elements of the node tree.
type BaseBranch struct {
	*BaseLeaf

	Childrens []Leaf
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// AdoptChild implements Branch interface.
func (b *BaseBranch) AdoptChild(child Leaf) {
	// Checking child ready to be adopted
	if child == nil ||
		child.Parent() != nil {
		log.Fatal("can't adopt the child. (child is nil or child parent is not nil)")
	}

	// Checking that child is not parent this node.
	var node Leaf = b
	for node.Parent() != nil {
		if node == child {
			log.Fatal("can't adopt child, child is an ancestor of node")
		}

		node = node.Parent()
	}

	// Adopting the child
	child.setParent(b)
	if b.Attached() {
		child.Attach(b.owner)
	}
}

// DropChild implements Branch interface.
func (b *BaseBranch) DropChild(child Leaf) {
	if child == nil ||
		child.Parent() == nil ||
		child.Attached() != b.Attached() {
		log.Fatal("can't drop the child. (child is nil or child parent is nil or child attach state is different)")
	}

	child.setParent(nil)
	if child.Attached() {
		child.Detach()
	}
}
