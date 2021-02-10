package painter

import "github.com/negrel/paon/internal/draw"

var _ draw.Painter = flex{}

type flex struct{}

func makeFlex() draw.Painter {
	return flex{}
}

func (f flex) Layout(ctx draw.Context) {}

func (f flex) Draw(ctx draw.Context) {}
