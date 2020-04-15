package widgets

import "github.com/negrel/ginger/v2/widgets/node"

// Layout are widgets that lay out is children.
type Layout interface {
	Widget
	node.Branch
}
