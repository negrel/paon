package tree

// ROOT is the root of the node tree
var ROOT *Layout = &Layout{
	BaseNode: &BaseNode{
		parent: nil,
	},
	childNodes: []Node{},
}
