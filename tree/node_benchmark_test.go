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

func newBenchIters(bN int, nbChild int) []benchIter {
	nodeData := time.Now()

	// In order to remove child multiple time, we must setup an array of parents.
	iters := make([]benchIter, bN)

	for i := 0; i < bN; i++ {
		parent := NewNode(nodeData)
		children := make([]*Node[time.Time], nbChild)

		for i := 0; i < nbChild; i++ {
			children[i] = NewNode(nodeData)
			err := parent.AppendChild(children[i])
			if err != nil {
				panic(err)
			}
		}

		rand.Shuffle(len(children), func(i, j int) {
			children[i], children[j] = children[j], children[i]
		})

		iters[i] = benchIter{parent, children}
	}

	return iters
}

func BenchmarkNodeAppendChild(b *testing.B) {
	for i := 64; i <= 512; i *= 2 {
		b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
			benchmarkNodeAppendChild(b, i)
		})
	}
}

func benchmarkNodeAppendChild(b *testing.B, nbChild int) {
	iters := newBenchIters(b.N, nbChild)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parent := iters[i].parent
		children := iters[i].children

		for _, child := range children {
			err := parent.AppendChild(child)
			if err != nil {
				panic(err)
			}
		}
	}
}

func BenchmarkNodeInsertBefore(b *testing.B) {
	for i := 64; i <= 512; i *= 2 {
		b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
			benchmarkNodeInsertBefore(b, i)
		})
	}
}

func benchmarkNodeInsertBefore(b *testing.B, nbChild int) {
	iters := newBenchIters(b.N, nbChild)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parent := iters[i].parent
		children := iters[i].children

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
	for i := 64; i <= 512; i *= 2 {
		b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
			benchmarkNodeRemoveChild(b, i)
		})
	}
}

func benchmarkNodeRemoveChild(b *testing.B, nbChild int) {
	iters := newBenchIters(b.N, nbChild)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parent := iters[i].parent
		children := iters[i].children

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
