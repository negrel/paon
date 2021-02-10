package render

import (
	"context"
	"github.com/negrel/debuggo/pkg/log"
	"time"
)

// Engine is responsible for rendering the
type Engine struct {
	ch     chan Buffer
	ctx    context.Context
	Screen Screen
	clock  *time.Ticker
}

// NewEngine return a new rendering engine that draw on the given surface.
func NewEngine(clock *time.Ticker, ctx context.Context) *Engine {
	return &Engine{
		ch:    make(chan Buffer),
		ctx:   ctx,
		clock: clock,
	}
}

func (e *Engine) Draw(patch Buffer) {
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
					e.Screen.Update()
					needRefresh = false
				}()
			}

		case <-e.ctx.Done():
			log.Debugln("stopping the rendering engine")
			return
		}
	}
}
