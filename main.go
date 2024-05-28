package main

import (
	"eng/game"
	"eng/sprite"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	Unit := 30
	init_X, init_Y := 225, 225

	sp_map := &sprite.MapAr{}
	sp_map.Load_file("sprite/gamemap.txt")

	Width, Height := sp_map.Size()
	Width *= Unit
	Height *= Unit
	fmt.Printf("Window Size: %v %v\n", Width, Height)

	sp := sprite.Sprites{
		X:      int(Width / 2),
		Y:      int(Height / 2),
		Unit:   Unit,
		C:      color.Black,
		Hitmap: sp_map,
		Plable: false,
	}

	player := sprite.Player(Unit - 1)
	player.X, player.Y = init_X, init_Y

	ebiten.SetWindowSize(Width, Height)
	ebiten.SetWindowTitle("Shaker JUMP")

	g := game.Game{
		ScreenHeight: Height,
		ScreenWidth:  Width,
		Unit:         Unit,
		Sprites:      []*sprite.Sprites{player, &sp},
	}

	ebiten.RunGame(&g)
}
