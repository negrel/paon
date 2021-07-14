package widgets

import (
	"fmt"
	"testing"
)

func BenchmarkWidget(b *testing.B) {
	b.Run("Root", benchmarksWidgetRoot)
}

func benchmarksWidgetRoot(b *testing.B) {
	for i := 8; i < 1024; i *= 2 {
		b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
			benchmarkWidgetRoot(b, i)
		})
	}
}

func benchmarkWidgetRoot(b *testing.B, deep int) {
	root := NewRoot()

	parent := newLayout()
	rootLayout := parent
	for i := 0; i < deep; i++ {
		child := newLayout()
		parent.AppendChild(child)
		parent = child
	}
	deepestChild := parent
	_ = deepestChild
	root.SetChild(rootLayout)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = deepestChild.Root()
	}
}
