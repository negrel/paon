package widgets

import "fmt"

var _ Layout = &SingleChildLayout{}

type SingleChildLayout struct {
	*Core
	child Widget
}

func (s SingleChildLayout) AppendChild(child Widget) {
	if s.child != nil {
		panic(fmt.Sprintf("%v can only have one child", s.Name()))
	}

	if child.Parent() != nil {
		panic(fmt.Sprintf("%v already have a parent", child.Name()))
	}

	s.child = child
	child.setParent(s)
}

func (s SingleChildLayout) IndexOf(child Widget) int {
	panic("implement me")
}

func (s SingleChildLayout) RemoveChild(child Widget) {
	panic("implement me")
}
