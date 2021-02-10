package painter

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/draw"
	"github.com/negrel/paon/pkg/pdk/styles/property"
)

// AddCustomRenderer adds the given draw.Painter constructor to the render map
// and return its ID. This ID should be used for the Display property.
func AddCustomRenderer(constructor func() draw.Painter) int {
	return rMap.add(constructor)
}

// GetFor returns a new draw.Painter for the given draw.Object.
func GetFor(object draw.Object) draw.Painter {
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
	renderers map[int]func() draw.Painter
}

const (
	Hidden = iota
	Block
	Inline
	Flex
)

var rMap = rendererMap{
	renderers: map[int]func() draw.Painter{
		Hidden: makeHidden,
		Block:  makeBlock,
		Inline: makeInline,
		Flex:   makeFlex,
	},
}

func (rm rendererMap) add(renderer func() draw.Painter) int {
	index := rm.len()
	rm.renderers[index] = renderer

	return index
}

func (rm rendererMap) len() int {
	return len(rm.renderers)
}
