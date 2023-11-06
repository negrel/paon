package draw

type Pipe func(Surface, Drawer)

// Pipeline define a drawing pipeline. It enable you to build non-coupled drawer
// easily. Most recently added pipes are executed first.
type Pipeline struct {
	pipe DrawerFunc
}

// NewPipeline returns a new pipeline based on the given layout.
func NewPipeline(drawer Drawer) Pipeline {
	return Pipeline{
		pipe: func(s Surface) {
			drawer.Draw(s)
		},
	}
}

// Draw implements Drawer.
func (p Pipeline) Draw(surface Surface) {
	p.pipe.Draw(surface)
}

// Pipe adds the given pipe to the pipeline.
func (p *Pipeline) Pipe(pipe Pipe) {
	next := p.pipe

	p.pipe = func(surface Surface) {
		pipe(surface, next)
	}
}
