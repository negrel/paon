//go:generate mockgen -destination mock/layout.go github.com/negrel/paon/internal/widgets Layout

package widgets

type Layout interface {
	Widget

	AppendChild(child Widget)
	IndexOf(child Widget) int
	RemoveChild(child Widget)
}

func newMockLayout() Layout {
	return nil
}
