package render

import (
	"context"
	"time"

	"github.com/negrel/debuggo/pkg/log"
)

// Engine is responsible for rendering the
type Engine struct {
	ch          chan Patch
	ctx         context.Context
	Screen      Screen
	ticker      *time.Ticker
	needRefresh bool
}

// NewEngine return a new rendering engine that draw on the given surface.
func NewEngine(screen Screen, ctx context.Context) *Engine {
	log.Infoln("creating the rendering engine")

	return &Engine{
		ch:     make(chan Patch),
		ctx:    ctx,
		Screen: screen,
		ticker: time.NewTicker(time.Millisecond * 16),
	}
}

func (e *Engine) Render(patch *Patch) {
	e.ch <- *patch
}

// SetUpdateInterval set the update interval of the surface.
// By default the rendering engine update the surface every 16ms (around 60fps).
func (e *Engine) SetUpdateInterval(duration time.Duration) {
	e.ticker = time.NewTicker(duration)
}

// Start the rendering engine.
func (e *Engine) Start() {
	log.Infoln("starting the rendering engine")

renderLoop:
	for {
		select {
		case patch := <-e.ch:
			e.Screen.Apply(patch)
			e.needRefresh = true

		case <-e.ticker.C:
			if e.needRefresh {
				e.Screen.Update()
				e.needRefresh = false
			}

		case <-e.ctx.Done():
			break renderLoop
		}
	}

	log.Infoln("stopping the rendering engine")
}
