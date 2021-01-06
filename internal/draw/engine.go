package draw

import (
	"context"
	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/pkg/runtime"
)

// Engine is responsible for rendering the
type Engine struct {
	ch          chan Canvas
	ctx         context.Context
	Screen      Screen
	needRefresh bool
}

// NewEngine return a new rendering engine that draw on the given surface.
func NewEngine(ctx context.Context) *Engine {
	return &Engine{
		ch:  make(chan Canvas),
		ctx: ctx,
	}
}

func (e *Engine) Draw(patch Canvas) {
	e.ch <- patch
}

// Start the rendering engine.
func (e *Engine) Start() {
	log.Debugln("starting the rendering engine")

renderLoop:
	for {
		select {
		case patch := <-e.ch:
			go func() {
				e.Screen.Apply(patch)
				e.needRefresh = true
			}()

		case <-runtime.Clock.C:
			if e.needRefresh {
				go func() {
					e.Screen.Update()
					e.needRefresh = false
				}()
			}

		case <-e.ctx.Done():
			break renderLoop
		}
	}

	log.Debugln("stopping the rendering engine")
}
