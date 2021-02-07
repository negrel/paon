package renderer

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/pkg/pdk/styles/property"
)

// AddCustomRenderer adds the given render.Renderer constructor to the render map
// and return its ID. This ID should be used for the Display property.
func AddCustomRenderer(constructor func() render.Renderer) int {
	return rMap.add(constructor)
}

// GetFor returns a new render.Renderer for the given render.Object.
func GetFor(object render.Object) render.Renderer {
	prop := object.Style().Get(property.IDDisplay)
	assert.NotNil(prop)
	rendererID := prop.(property.Int)

	return rMap.renderers[rendererID.Value]()
}

// IsValidRendererID returns true if the given rendererID is valid.
func IsValidRendererID(rendererID int) bool {
	return rendererID > 0 && rendererID < rMap.len()
}

type rendererMap struct {
	renderers map[int]func() render.Renderer
}

const (
	Hidden = iota
	Block
	Inline
	Flex
)

var rMap = rendererMap{
	renderers: map[int]func() render.Renderer{
		Hidden: makeHidden,
		Block:  makeBlock,
		Inline: makeInline,
		Flex:   makeFlex,
	},
}

func (rm rendererMap) add(renderer func() render.Renderer) int {
	index := rm.len()
	rm.renderers[index] = renderer

	return index
}

func (rm rendererMap) len() int {
	return len(rm.renderers)
}
