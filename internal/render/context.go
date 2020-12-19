package render

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/draw"
	"github.com/negrel/paon/internal/geometry"
)

// Context is the rendering context passed to Object for the layout step and the draw step.
type Context struct {
	Object Object
	canvas *draw.Patch
}

func MakeContext(obj Object, canvas *draw.Patch) Context {
	return Context{
		Object: obj,
		canvas: canvas,
	}
}

// Layer return a geometry.Rectangle object defining the current Object layer.
func (ctx Context) Layer() *geometry.Rectangle {
	return &ctx.canvas.Rectangle
}

// Canvas return the draw.Patch of the rendering context.
// A mock canvas is returned during LayoutStep.
func (ctx Context) Canvas() draw.Patch {
	assert.Equal(ctx.Step(), DrawStepType, "accessing canvas during non-draw step is not allowed")

	if ctx.Step() == DrawStepType {
		return *ctx.canvas
	}

	return draw.Patch{}
}

func (ctx Context) Step() StepType {
	return ctx.Object.RenderStep()
}
