package renderer

import (
	"github.com/negrel/paon/internal/render"
)

type RenderMap struct {
	renderers map[int]func() render.Renderer
}

const (
	Hidden = iota
	Block
	Flex
)

var Map = RenderMap{
	renderers: map[int]func() render.Renderer{
		Hidden: makeHidden,
		Block:  makeBlock,
		Flex:   makeFlex,
	},
}

func (rm RenderMap) Add(renderer func() render.Renderer) int {
	index := rm.Len()
	rm.renderers[index] = renderer

	return index
}

func (rm RenderMap) Len() int {
	return len(rm.renderers)
}
