//go:generate mockgen -source ./widget.go -destination mock/widget.go Widget

package widgets

type Widget interface {
	Name() string

	Parent() Layout
	setParent(parent Layout)

	build(ctx interface{})
}
