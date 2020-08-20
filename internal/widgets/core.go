package widgets

var _ Widget = &Core{}

// Core define the core of every widget.
type Core struct {
	name   string
	parent Layout
}

func newCore(name string) *Core {
	return &Core{
		name: name,
	}
}

func (c *Core) build(ctx interface{}) {
	panic("implement me")
}

func (c *Core) Name() string {
	return c.name
}

func (c *Core) Parent() Layout {
	return c.parent
}

func (c *Core) setParent(parent Layout) {
	c.parent = parent
}
