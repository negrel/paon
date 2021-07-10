package tree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkNodeAppendChild(b *testing.B) {
	parent := NewNode(nodeData)
	nodes := make([]Node, b.N)

	for i := 0; i < b.N; i++ {
		nodes[i] = newLeafNode(nodeData)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = parent.AppendChild(nodes[i])
	}
}

func BenchmarkNodeInsertBefore(b *testing.B) {
	parent := NewNode(nodeData)
	nodes := make([]Node, b.N)

	for i := 0; i < b.N; i++ {
		nodes[i] = newLeafNode(nodeData)
	}

	var previousNode Node
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = parent.InsertBefore(previousNode, nodes[i])
		previousNode = nodes[i]
	}
}

func BenchmarkNodeRemoveChild(b *testing.B) {
	parent := NewNode(nodeData)
	nodes := make([]Node, b.N)

	for i := 0; i < b.N; i++ {
		nodes[i] = newLeafNode(nodeData)
		parent.AppendChild(nodes[i])
	}

	rand.Shuffle(len(nodes), func(i, j int) {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = parent.RemoveChild(nodes[i])
	}
}

func BenchmarkNodeRoot(b *testing.B) {
	b.Run("Root", func(b *testing.B) {
		for i := 8; i < 1024; i *= 2 {
			b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
				benchmarkNodeRoot(b, i)
			})
		}
	})
}

func benchmarkNodeRoot(b *testing.B, deep int) {
	rootData := nodeData.Add(time.Minute)
	root := NewRoot(rootData)

	var parent Node = root
	for i := 0; i < deep; i++ {
		child := NewNode(nodeData)
		parent.AppendChild(child)
		parent = child
	}
	deepestChild := parent

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = deepestChild.Root()
	}
}
