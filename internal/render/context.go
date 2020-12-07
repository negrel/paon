package render

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/draw"
	"github.com/negrel/paon/internal/geometry"
)

// Context is the rendering context passed to Object for the layout step and the draw step.
type Context struct {
	step StepType

	Object Object
	canvas *draw.Canvas
}

// Layer return a geometry.Rectangle object defining the current Object layer.
func (ctx Context) Layer() *geometry.Rectangle {
	return &ctx.canvas.Rectangle
}

// Canvas return the draw.Canvas of the rendering context.
// A mock canvas is returned during LayoutStep.
func (ctx Context) Canvas() draw.Canvas {
	assert.Equal(ctx.step, DrawStepType, "accessing canvas during non-draw step is not allowed")

	if ctx.step == DrawStepType {
		return *ctx.canvas
	}

	return draw.Canvas{}
}

func (ctx Context) Step() StepType {
	return ctx.step
}
