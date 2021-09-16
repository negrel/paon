package main

import (
	"image"
	"time"

	"github.com/nfnt/resize"

	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/layout"
	pdkwidgets "github.com/negrel/paon/pdk/widgets"
	"github.com/negrel/paon/styles/value"
)

type imv struct {
	*pdkwidgets.BaseWidget
	src image.Image
	rsz image.Image
}

func NewIMV(img image.Image) *imv {
	imv := &imv{
		src: img,
	}
	imv.BaseWidget = pdkwidgets.NewBaseWidget(
		pdkwidgets.LayoutManager(layout.ManagerFn(func(c layout.Constraint) layout.BoxModel {
			imv.rsz = resize.Resize(uint(c.Max.Width()), uint(c.Max.Height()), imv.src, resize.NearestNeighbor)

			return layout.NewBox(c.Max)
		})),
		pdkwidgets.Drawer(draw.DrawerFn(func(ctx *draw.Context) {
			ctx.AddOp(func(c draw.Canvas) {
				start := time.Now()
				bounds := imv.Box().ContentBox()

				for i := bounds.Min.X(); i < bounds.Max.X(); i++ {
					for j := bounds.Min.Y(); j < bounds.Max.Y(); j++ {
						r, g, b, _ := imv.rsz.At(i, j).RGBA()

						c.Set(geometry.Pt(i, j), draw.Cell{
							Content: '\u2588',
							Style: draw.CellStyle{
								Foreground: value.ColorFromRGB(uint8(r>>8), uint8(g>>8), uint8(b>>8)),
							},
						})
					}
				}
				log.Debug(time.Now().Sub(start))
			})
		})),
	)

	return imv
}
