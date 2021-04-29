package tree

type Root struct {
	*parentNode
}

func NewRoot(data interface{}) Node {
	return Root{
		parentNode: newNode(data),
	}
}

func (r Root) Root() Node {
	return r
}

func (r Root) Parent() Node {
	return nil
}

func (r Root) AppendChild(newChild Node) error {
	err := r.parentNode.AppendChild(newChild)
	if err != nil {
		return err
	}
	newChild.SetParent(r)

	return nil
}

func (r Root) InsertBefore(reference, newChild Node) error {
	err := r.parentNode.InsertBefore(reference, newChild)
	if err != nil {
		return err
	}
	newChild.SetParent(r)

	return nil
}
