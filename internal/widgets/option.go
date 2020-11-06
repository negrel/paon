package widgets

import (
	"github.com/negrel/paon/internal/render"
)

type Option struct {
	priority int
	apply    func(render.Surface)
}

func Opt(apply func(render.Surface), priority int) Option {
	return Option{
		priority: priority,
		apply:    apply,
	}
}
