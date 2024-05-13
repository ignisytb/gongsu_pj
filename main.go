package main

import (
	"eng/game"
	"eng/sprite"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	Height := 1050
	Width := 450

	ebiten.SetWindowSize(Width, Height)
	ebiten.SetWindowTitle("Shaker JUMP")

	shaker := sprite.Shaker(float64(Width/2), float64(Height/2), 0)

	g := game.Game{
		ScreenHeight: Height,
		ScreenWidth:  Width,
		Sprites:      []*sprite.Sprites{shaker},
	}

	ebiten.RunGame(&g)
}
