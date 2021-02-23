package gfx

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mrherbivour/StrongAuras/config"
)

type Gfx struct {
	Config *config.Config
}

var rMouseDown bool

func (g *Gfx) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		rMouseDown = true
	} else if rMouseDown {
		ToggleMovable()
		rMouseDown = false
	}

	return nil
}

func (g *Gfx) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{
		0, 0, 0, 80,
	})
	for _, ind := range g.Config.Indicators {
		if ind.EImg != nil {
			if ind.Visable || ind.Duration > 0 {
				op := &ebiten.DrawImageOptions{}
				if ind.X > 0 || ind.Y > 0 {
					op.GeoM.Translate(ind.X, ind.Y)
				}
				eImgSize := ind.EImg.Bounds().Size()
				if ind.W > 0 {
					op.GeoM.Scale(ind.W/float64(eImgSize.X), 1)
				}
				if ind.H > 0 {
					op.GeoM.Scale(1, ind.H/float64(eImgSize.Y))
				}
				if !ind.Visable {
					sec := time.Now().Sub(ind.HiddenAt).Seconds()
					scale := sec / float64(ind.Duration)
					if scale > 1 {
						scale = 1
					}
					op.GeoM.Scale(scale, scale)
					if scale < 1 {
						op.ColorM.Scale(1, 1, 1, 0.25+scale/2)
					}
					op.ColorM.Apply(color.Transparent)
				}
				screen.DrawImage(ind.EImg, op)
			}
		}
	}
}

func (g *Gfx) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Config.WindowPosition.W, g.Config.WindowPosition.H
}
