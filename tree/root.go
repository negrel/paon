package tree

import "github.com/negrel/paon/events"

// IsRoot returns whether a node is a root node.
func IsRoot[T events.Target](n Node[T]) bool {
	return n != nil && n.RootNode() == n
}

type root[T events.Target] struct {
	*node[T]
}

// NewRoot creates a new root node. Node are considered as mounted only if
// one of their ancestor is a root node.
func NewRoot[T events.Target](data T) Node[T] {
	return root[T]{
		node: newNode[T](data),
	}
}

// RootNode implements Node.
func (r root[T]) RootNode() Node[T] {
	return r
}
