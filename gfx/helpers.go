package gfx

import "github.com/hajimehoshi/ebiten/v2"

func ToggleMovable() {
	ebiten.SetWindowDecorated(
		!ebiten.IsWindowDecorated(),
	)
}
