package layout

import "github.com/negrel/paon/geometry"

type Pipe func(Constraint, Layout) geometry.Size

// Pipeline define a layout pipeline. It enable you to build non-coupled layout
// easily. Most recently added pipes are executed first.
type Pipeline struct {
	pipe LayoutFunc
}

// NewPipeline returns a new pipeline based on the given layout.
func NewPipeline(layout Layout) Pipeline {
	return Pipeline{
		pipe: func(co Constraint) geometry.Size {
			return layout.Layout(co)
		},
	}
}

// Layout implements Layout.
func (p Pipeline) Layout(co Constraint) geometry.Size {
	return p.pipe(co)
}

// Pipe adds the given pipe to the pipeline.
func (p *Pipeline) Pipe(pipe Pipe) {
	next := p.pipe

	p.pipe = func(co Constraint) geometry.Size {
		return pipe(co, next)
	}
}
