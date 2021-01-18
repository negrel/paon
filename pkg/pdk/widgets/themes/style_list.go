package themes

import (
	"github.com/negrel/debuggo/pkg/assert"
	pdkstyles "github.com/negrel/paon/pkg/pdk/styles"
)

type styleNode struct {
	pdkstyles.Style
	next *styleNode
}

type styleList struct {
	first *styleNode
}

func (l *styleList) append(s pdkstyles.Style) {
	assert.NotNil(s)

	node := &styleNode{
		Style: s,
		next:  nil,
	}

	if l.first == nil {
		l.first = node
		return
	}

	n := l.first
	for n.next != nil {
		n = n.next
	}

	n.next = node
}

func (l *styleList) remove(s pdkstyles.Style) {
	assert.NotNil(s)

	node := &l.first

	for (*node) != nil {
		if (*node).Style == s {
			*node = (*node).next
			break
		}

		node = &(*node).next
	}
}

func (l *styleList) values() []pdkstyles.Style {
	var result []pdkstyles.Style

	node := l.first
	for node != nil {
		result = append(result, node.Style)
		node = node.next
	}

	return result
}
