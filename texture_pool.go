package main

import (
	"fmt"
	"image"
	"image/draw"

	"github.com/gonutz/d3d9"
)

// type saTexture struct {
// 	D3D9Texture *d3d9.Texture
// }

var texturePool map[string]*d3d9.Texture

func loadTexture(filename string) {
	if texturePool == nil {
		texturePool = make(map[string]*d3d9.Texture)
	}
	if texturePool[filename] != nil {
		return
	}
	img := loadSprite(
		fmt.Sprint(cfg.EqFolder, "uifiles\\", filename),
	)

	var nrgba *image.NRGBA
	if asRGBA, ok := img.(*image.NRGBA); ok {
		nrgba = asRGBA
	} else {
		fmt.Println("!!!! Well shit !!!")
		nrgba = image.NewNRGBA(img.Bounds())
		draw.Draw(nrgba, nrgba.Bounds(), img, image.ZP, draw.Src)
	}
	texture, err := device.CreateTexture(
		uint(nrgba.Bounds().Dx()),
		uint(nrgba.Bounds().Dy()),
		1,
		d3d9.USAGE_SOFTWAREPROCESSING,
		d3d9.FMT_A8R8G8B8,
		d3d9.POOL_MANAGED,
		0,
	)
	check(err)
	lockedRect, err := texture.LockRect(0, nil, d3d9.LOCK_DISCARD)
	check(err)
	b := nrgba.Pix
	// For some reason Go image and DirectX texture are in an opposite order for colors here
	// When RGBA is here from go, we need to send it as BGRA.
	for i := 0; i < len(b)/4; i++ {
		l := i * 4
		b[l], b[l+2] = b[l+2], b[l]
	}
	lockedRect.SetAllBytes(b, nrgba.Stride)
	check(texture.UnlockRect(0))
	texturePool[filename] = texture
}
