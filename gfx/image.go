package gfx

import (
	"os"

	"github.com/ftrvxmtrx/tga"
	"github.com/hajimehoshi/ebiten/v2"
)

func LoadSprite(filename string) *ebiten.Image {
	if spriteCache[filename] != nil {
		return spriteCache[filename]
	}
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, err := tga.Decode(f)
	if err != nil {
		panic(err)
	}

	spriteCache[filename] = ebiten.NewImageFromImage(img)
	return spriteCache[filename]
}
