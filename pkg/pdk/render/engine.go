package render

import (
	"context"
	"github.com/negrel/debuggo/pkg/log"
	"time"
)

// Engine is responsible for rendering the received buffers.
type Engine struct {
	ch     chan Patch
	ctx    context.Context
	Screen Surface
	clock  *time.Ticker
}

// NewEngine return a new rendering engine that draw on the given surface.
func NewEngine(clock *time.Ticker, ctx context.Context) *Engine {
	return &Engine{
		ch:    make(chan Patch),
		ctx:   ctx,
		clock: clock,
	}
}

func (e *Engine) Draw(patch Patch) {
	e.ch <- patch
}

// Start the rendering engine.
func (e *Engine) Start() {
	log.Debugln("starting the rendering engine")

	needRefresh := false

	for {
		select {
		case patch := <-e.ch:
			go func() {
				e.Screen.Apply(patch)
				needRefresh = true
			}()

		case <-e.clock.C:
			if needRefresh {
				go func() {
					e.Screen.Flush()
					needRefresh = false
				}()
			}

		case <-e.ctx.Done():
			log.Debugln("stopping the rendering engine")
			return
		}
	}
}
