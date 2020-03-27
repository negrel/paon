package tree

// NodeList objects are collections of nodes (live)
// https://developer.mozilla.org/en-US/docs/Web/API/NodeList
type NodeList struct {
	list []*Node
}

func newNodeList() *NodeList {
	return &NodeList{}
}

func (nl *NodeList) append(node *Node) Node {
	var child = node
	nl.list = append(nl.list, child)
	return *child
}

func (nl *NodeList) appendList(nodes ...*Node) {
	nl.list = append(nl.list, nodes...)
}

func (nl *NodeList) set(index int, node *Node) {
	nl.list[index] = node
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/
// ANCHOR Getters & Setters

// Length method return the number of node in the list
func (nl *NodeList) Length() int {
	return len(nl.Values())
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// ForEach apply the given function for each of
// the Node in the list.
func (nl *NodeList) ForEach(fn func(i int, c *Node)) {
	for i, v := range nl.Values() {
		fn(i, v)
	}
}

// IndexOf method return the index of the
// searched node and return -1 if not found.
func (nl *NodeList) IndexOf(searched *Node) int {
	for index, node := range nl.Values() {
		if same := node.IsSameNode(searched); same {
			return index
		}
	}

	return -1
}

// Item return a node from the Node list by index
func (nl *NodeList) Item(index int) *Node {
	if index >= 0 && index < nl.Length() {
		return nl.Values()[index]
	}
	return nil
}

// Values method returns an iterator allowing to go
// through all values contained in this object.
// https://developer.mozilla.org/en-US/docs/Web/API/NodeList/values
func (nl *NodeList) Values() []*Node {
	return nl.list
}
