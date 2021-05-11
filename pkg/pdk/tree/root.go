package tree

type Root struct {
	*node
}

func NewRoot(data interface{}) Node {
	return Root{
		node: newNode(data),
	}
}

func (r Root) Root() Node {
	return r
}

func (r Root) Parent() Node {
	return nil
}

func (r Root) AppendChild(newChild Node) error {
	err := r.node.AppendChild(newChild)
	if err != nil {
		return err
	}
	newChild.SetParent(r)

	return nil
}

func (r Root) InsertBefore(reference, newChild Node) error {
	err := r.node.InsertBefore(reference, newChild)
	if err != nil {
		return err
	}
	newChild.SetParent(r)

	return nil
}
