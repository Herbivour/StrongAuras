package main

import (
	"image"
	"os"

	"github.com/ftrvxmtrx/tga"
	"github.com/gonutz/d3d9"
)

var t *d3d9.Texture

func loadSprite(filename string) image.Image {
	f, err := os.Open(filename)
	defer f.Close()
	check(err)
	img, err := tga.Decode(f)
	check(err)

	return img
}
