package events

import (
	"github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/pdk/tree"
)

type Option func(*node)

// InternalNode return an Option that sets the internal tree.Node used by the Node.
func InternalNode(in tree.Node) Option {
	return func(n *node) {
		n.node = in
	}
}

// Target returns an Option that sets the internal events.Target used by the Node.
func Target(target events.Target) Option {
	return func(n *node) {
		n.Target = target
	}
}
