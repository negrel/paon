package tree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type benchIter struct {
	parent   *Node[time.Time]
	children []*Node[time.Time]
}

func newBenchIter(nbChild int, appendChild bool) benchIter {
	nodeData := time.Now()

	// In order to remove child multiple time, we must setup an array of parents.
	iter := benchIter{
		parent:   NewNode(nodeData),
		children: make([]*Node[time.Time], nbChild),
	}

	for i := 0; i < nbChild; i++ {
		iter.children[i] = NewNode(nodeData)

		if appendChild {
			err := iter.parent.AppendChild(iter.children[i])
			if err != nil {
				panic(err)
			}
		}
	}

	if appendChild {
		rand.Shuffle(len(iter.children), func(i, j int) {
			iter.children[i], iter.children[j] = iter.children[j], iter.children[i]
		})
	}

	return iter
}

func BenchmarkNodeAppendChild(b *testing.B) {
	for i := 64; i <= 128; i *= 2 {
		b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
			benchmarkNodeAppendChild(b, i)
		})
	}
}

func benchmarkNodeAppendChild(b *testing.B, nbChild int) {
	for i := 0; i < b.N; i++ {
		iter := newBenchIter(nbChild, true)
		parent := iter.parent
		children := iter.children

		for _, child := range children {
			err := parent.AppendChild(child)
			if err != nil {
				panic(err)
			}
		}
	}
}

func BenchmarkNodeInsertBefore(b *testing.B) {
	for i := 64; i <= 128; i *= 2 {
		b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
			benchmarkNodeInsertBefore(b, i)
		})
	}
}

func benchmarkNodeInsertBefore(b *testing.B, nbChild int) {
	for i := 0; i < b.N; i++ {
		iter := newBenchIter(nbChild, true)
		parent := iter.parent
		children := iter.children

		var previousNode *Node[time.Time]

		for j, child := range children {
			err := parent.InsertBefore(child, previousNode)
			if err != nil {
				panic(err)
			}

			if j%2 == 0 {
				previousNode = child
			}
		}
	}
}

func BenchmarkNodeRemoveChild(b *testing.B) {
	for i := 64; i <= 128; i *= 2 {
		b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
			benchmarkNodeRemoveChild(b, i)
		})
	}
}

func benchmarkNodeRemoveChild(b *testing.B, nbChild int) {
	for i := 0; i < b.N; i++ {
		iter := newBenchIter(nbChild, true)
		parent := iter.parent
		children := iter.children

		for _, child := range children {
			err := parent.RemoveChild(child)
			if err != nil {
				panic(err)
			}
		}
	}
}

func BenchmarkNodeRoot(b *testing.B) {
	for i := 64; i <= 4096; i *= 2 {
		b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
			benchmarkNodeRoot(b, i)
		})
	}
}

func benchmarkNodeRoot(b *testing.B, deep int) {
	nodeData := time.Now()

	rootData := nodeData.Add(time.Minute)
	root := NewNode(rootData)

	var parent *Node[time.Time] = root
	for i := 0; i < deep; i++ {
		child := NewNode(nodeData)
		err := parent.AppendChild(child)
		if err != nil {
			panic(err)
		}
		parent = child
	}
	deepestChild := parent

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := deepestChild.Root()
		if r == deepestChild { // Should never happen.
			i++
		}
	}
}
