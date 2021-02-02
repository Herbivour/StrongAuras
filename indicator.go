package main

import (
	"strings"
)

type spritBox struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	W float32 `json:"w"`
	H float32 `json:"h"`
}
type indicator struct {
	Name           string   `json:"name"`
	Visable        bool     `json:"default"`
	ShowWhen       *string  `json:"show_when,omitempty"`
	HideWhen       *string  `json:"hide_when,omitempty"`
	SpriteSheet    *string  `json:"sprite_sheet"`
	SpriteSheetBox spritBox `json:"sprite_box"`
	X              float32  `json:"x"`
	Y              float32  `json:"y"`
	W              float32  `json:"w"`
	H              float32  `json:"h"`

	Triangles []vertex
	Id        int
}

func (ind *indicator) ProcessLogLine(text string) {
	if ind.HideWhen != nil && *ind.HideWhen != "" && strings.Contains(text, *ind.HideWhen) {
		ind.Visable = false
	} else if ind.ShowWhen != nil && *ind.ShowWhen != "" && strings.Contains(text, *ind.ShowWhen) {
		ind.Visable = true
	}
}

func (ind *indicator) InitTexture() {
	if ind.SpriteSheet != nil && *ind.SpriteSheet != "" {
		loadTexture(*ind.SpriteSheet)
	}
	left := ind.SpriteSheetBox.X
	right := left + ind.SpriteSheetBox.W
	top := ind.SpriteSheetBox.Y
	bottom := top + ind.SpriteSheetBox.H

	ind.Triangles = []vertex{
		{ind.X, ind.Y, right, top},
		{ind.X, ind.Y - ind.H, right, bottom},
		{ind.X - ind.W, ind.Y - ind.H, left, bottom},
		{ind.X - ind.W, ind.Y - ind.H, left, bottom},
		{ind.X - ind.W, ind.Y, left, top},
		{ind.X, ind.Y, right, top},
	}
}

func initIndicators() {
	// Prepare our icons
	for i, ind := range cfg.Indicators {
		ind.Id = i
		ind.InitTexture()
	}
	buildVertBuff()
}
