package config

import (
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type spritBox struct {
	X1 int `json:"x1"`
	Y1 int `json:"y1"`
	X2 int `json:"x2"`
	Y2 int `json:"y2"`
}
type indicator struct {
	Name           string    `json:"name"`
	Visable        bool      `json:"default"`
	ShowWhen       *string   `json:"show_when,omitempty"`
	HideWhen       *string   `json:"hide_when,omitempty"`
	Duration       int       `json:"duration"`
	SpriteSheet    *string   `json:"sprite_sheet"`
	SpriteSheetBox *spritBox `json:"sprite_box"`
	X              float64   `json:"x"`
	Y              float64   `json:"y"`
	W              float64   `json:"w"`
	H              float64   `json:"h"`
	Id             int

	EImg     *ebiten.Image `json:"-"`
	HiddenAt time.Time     `json:"-"`
}

func (ind *indicator) ProcessLogLine(text string) {
	if ind.HideWhen != nil && *ind.HideWhen != "" && strings.Contains(text, *ind.HideWhen) {
		ind.Visable = false
		ind.HiddenAt = time.Now()
	} else if ind.ShowWhen != nil && *ind.ShowWhen != "" && strings.Contains(text, *ind.ShowWhen) {
		ind.Visable = true
	}
}
