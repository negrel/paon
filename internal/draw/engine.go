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
	needRefresh bool
}

// NewEngine return a new rendering engine that draw on the given surface.
func NewEngine(screen Window, ctx context.Context) *Engine {
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
				runtime.Window.Apply(patch)
				e.needRefresh = true
			}()

		case <-runtime.Clock.C:
			if e.needRefresh {
				go func() {
					runtime.Window.Update()
					e.needRefresh = false
				}()
			}

		case <-e.ctx.Done():
			break renderLoop
		}
	}

	log.Debugln("stopping the rendering engine")
}
