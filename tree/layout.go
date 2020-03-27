package tree

// Layout is node that can contain other node (Layour OR View).
type Layout struct {
	*Node
	childNodes *NodeList
	style      Style
}

func newLayout() *Layout {
	return &Layout{
		&Node{
			parent: nil,
			id: "",
			class: []string{},
			isConnected: false,
		}
	}
}
