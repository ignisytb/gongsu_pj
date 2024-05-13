package main

import (
	"eng/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	Height := 480
	Width := 640

	ebiten.SetWindowSize(Width, Height)
	ebiten.SetWindowTitle("Shaker JUMP")

	g := game.Game{
		ScreenHeight: Height,
		ScreenWidth:  Width,
	}

	ebiten.RunGame(&g)
}
