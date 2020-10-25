//go:generate mockgen -destination mock/layout.go github.com/negrel/paon/internal/widgets Layout

package widgets

type Layout interface {
	Widget

	Append(child Widget)
	InsertBefore(child, reference Widget)
	Insert(child Widget, index int)
	Children() []Widget
	IndexOf(child Widget) int
	Remove(child Widget)
}
