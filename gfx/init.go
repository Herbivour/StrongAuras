package gfx

import (
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mrherbivour/StrongAuras/config"
)

var spriteCache map[string]*ebiten.Image

// Init - Inits the ebiten graphics window
func Init(cfg *config.Config) {
	spriteCache = make(map[string]*ebiten.Image)
	ebiten.SetInitFocused(false)
	ebiten.SetWindowSize(cfg.WindowPosition.W, cfg.WindowPosition.H)
	ebiten.SetWindowTitle("Strong Auras")
	ebiten.SetScreenTransparent(true)
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowPosition(cfg.WindowPosition.X, cfg.WindowPosition.Y)
	ebiten.SetWindowResizable(true)
	g := Gfx{
		cfg,
	}

	for _, ind := range cfg.Indicators {
		if ind.SpriteSheet != nil {
			spriteSrc := LoadSprite(
				fmt.Sprint(cfg.EqFolder, "uifiles\\", *ind.SpriteSheet),
			)
			if ind.SpriteSheetBox != nil {
				fmt.Println("Rect:", image.Rect(
					ind.SpriteSheetBox.X1,
					ind.SpriteSheetBox.Y1,
					ind.SpriteSheetBox.X2,
					ind.SpriteSheetBox.Y2,
				))
				ind.EImg = spriteSrc.SubImage(
					image.Rect(
						ind.SpriteSheetBox.X1,
						ind.SpriteSheetBox.Y1,
						ind.SpriteSheetBox.X2,
						ind.SpriteSheetBox.Y2,
					),
				).(*ebiten.Image)
				fmt.Println(ind.EImg)
			} else {
				ind.EImg = spriteSrc
			}
		}
	}

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
