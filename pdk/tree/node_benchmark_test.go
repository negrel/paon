package tree

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkNode_AppendChild(b *testing.B) {
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

func BenchmarkNode_InsertBefore(b *testing.B) {
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

func BenchmarkNode_RemoveChild(b *testing.B) {
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

func benchmarkNode_Root(b *testing.B, deep int) {
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

func BenchmarkNode_Root8(b *testing.B) {
	benchmarkNode_Root(b, 8)
}

func BenchmarkNode_Root16(b *testing.B) {
	benchmarkNode_Root(b, 16)
}
func BenchmarkNode_Root32(b *testing.B) {
	benchmarkNode_Root(b, 32)
}

func BenchmarkNode_Root64(b *testing.B) {
	benchmarkNode_Root(b, 64)
}

func BenchmarkNode_Root128(b *testing.B) {
	benchmarkNode_Root(b, 128)
}

func BenchmarkNode_Root256(b *testing.B) {
	benchmarkNode_Root(b, 256)
}

func BenchmarkNode_Root512(b *testing.B) {
	benchmarkNode_Root(b, 512)
}

func BenchmarkNode_Root1024(b *testing.B) {
	benchmarkNode_Root(b, 1024)
}
