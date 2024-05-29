package main

import (
	"eng/game"
	"eng/sprite"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	unit, n_th, n_sp := 30, 2, 4
	po_x, po_y := 285, 225
	th_x, th_y := []float64{15, 555}, []float64{285, 285}
	sp_x, sp_y := []float64{46, 524, 46, 524}, []float64{46, 46, 524, 524}

	gmp := &sprite.Sprites{Unit: unit, C: color.Black, Hitmap: &sprite.MapAr{}}
	gmp.Hitmap.Load_file("sprite/gamemap.txt")

	width, height := gmp.Hitmap.Size()
	gmp.X, gmp.Y = float64(width*unit)/2, float64(height*unit)/2

	ebiten.SetWindowSize(width*unit, height*unit)
	ebiten.SetWindowTitle("PPAP MAN")

	g := &game.Game{
		GameOver:     false,
		ScreenWidth:  width * unit,
		ScreenHeight: height * unit,
		Unit:         unit,
		Current_Ply:  1,
		Num_Thief:    n_th,
		Num_Sp:       n_sp,
		Sprites:      []*sprite.Sprites{gmp},
	}

	police := sprite.Player(unit - 1)
	police.X = float64(po_x)
	police.Y = float64(po_y)
	police.C = color.RGBA{0, 0, 255, 100}
	police.CharSpd = 1.5
	g.Sprites = append(g.Sprites, police)

	for i := 0; i < n_th; i++ {
		th := sprite.Player(unit - 1)
		th.X = th_x[i]
		th.Y = th_y[i]
		th.C = color.RGBA{255, 0, 0, 200}
		g.Sprites = append(g.Sprites, th)
	}

	for i := 0; i < n_sp; i++ {
		sp := sprite.Player(unit - 2)
		sp.X = sp_x[i]
		sp.Y = sp_y[i]
		sp.C = color.RGBA{100, 100, 100, 200}
		g.Sprites = append(g.Sprites, sp)
	}

	ebiten.RunGame(g)
}
