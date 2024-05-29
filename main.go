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

	gmp := &sprite.Sprites{Motion: &sprite.Motion{}, Unit: unit, C: color.Black, Hitmap: &sprite.MapAr{}}
	gmp.Hitmap.Load_file("sprite/gamemap.txt")

	width, height := gmp.Hitmap.Size()
	gmp.Motion.X, gmp.Motion.Y = float64(width*unit)/2, float64(height*unit)/2

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
	police.Motion.X = float64(po_x)
	police.Motion.Y = float64(po_y)
	police.C = color.RGBA{0, 0, 255, 100}
	police.Motion.MaxSpd = 1.5
	g.Sprites = append(g.Sprites, police)

	for i := 0; i < n_th; i++ {
		th := sprite.Player(unit - 1)
		th.Motion.X = th_x[i]
		th.Motion.Y = th_y[i]
		th.Motion.MaxSpd = 2
		th.C = color.RGBA{255, 0, 0, 200}
		g.Sprites = append(g.Sprites, th)
	}

	for i := 0; i < n_sp; i++ {
		sp := sprite.Player(unit - 2)
		sp.Motion.X = sp_x[i]
		sp.Motion.Y = sp_y[i]
		sp.C = color.RGBA{100, 100, 100, 200}
		g.Sprites = append(g.Sprites, sp)
	}

	ebiten.RunGame(g)
}
