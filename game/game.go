package game

import (
	"eng/sprite"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ScreenHeight int
	ScreenWidth  int
	Unit         int
	Current_Ply  int
	Num_Thief    int
	Num_Sp       int
	Sprites      []*sprite.Sprites
	GameOver     bool
}

func (g *Game) Update() error {
	if !g.GameOver {
		wall := g.Sprites[0]
		center_x, center_y := float64(g.ScreenWidth/2), float64(g.ScreenHeight/2)

		g.Sprites[1].Police_Move(wall)
		if ebiten.IsKeyPressed(ebiten.Key1) {
			g.Current_Ply = 1
		} else if ebiten.IsKeyPressed(ebiten.Key2) {
			g.Current_Ply = 2
		}

		save_offset := 2
		if !g.Sprites[g.Current_Ply+1].Thief_Move(wall) {
			if !(g.Sprites[g.Current_Ply+1].Motion.X == center_x && g.Sprites[g.Current_Ply+1].Motion.Y == center_y) {
				if !(g.Sprites[g.Current_Ply+1].Motion.X < (center_x+(float64(2*g.Unit)+float64(save_offset))) &&
					g.Sprites[g.Current_Ply+1].Motion.X > (center_x-(float64(2*g.Unit)+float64(save_offset))) &&
					g.Sprites[g.Current_Ply+1].Motion.Y < (center_y+(float64(2*g.Unit)+float64(save_offset))) &&
					g.Sprites[g.Current_Ply+1].Motion.Y > (center_y-(float64(2*g.Unit)+float64(save_offset)))) {
					for i := 0; i < g.Num_Thief; i++ {
						if g.Sprites[i+2].Motion.X == center_x && g.Sprites[i+2].Motion.Y == center_y {
							g.Sprites[i+2].Jail_Ent = time.Now()
						}
					}
				}
			} else {
				for i := 0; i < g.Num_Thief; i++ {
					if g.Sprites[i+2].Motion.X == center_x && g.Sprites[i+2].Motion.Y == center_y {
						g.Sprites[i+2].Jail_Ent = time.Now()
					}
				}
			}
		} else {
			for i := 0; i < g.Num_Thief; i++ {
				if g.Sprites[i+2].Motion.X == center_x && g.Sprites[i+2].Motion.Y == center_y {
					g.Sprites[i+2].Jail_Ent = time.Now()
				}
			}
		}

		for i := 0; i < g.Num_Thief; i++ {
			if g.Sprites[1].Coll(g.Sprites[i+2]) {
				g.Sprites[i+2].Motion.X = center_x
				g.Sprites[i+2].Motion.Y = center_y
				g.Sprites[i+2].Jail_Ent = time.Now()
			}
		}

		for i := 0; i < g.Num_Thief; i++ {
			if g.Sprites[i+2].Motion.X == center_x && g.Sprites[i+2].Motion.Y == center_y {
				if time.Since(g.Sprites[i+2].Jail_Ent).Seconds() >= 2.0 {
					g.Sprites[i+2].Motion.X = center_x
					g.Sprites[i+2].Motion.Y = center_y - float64(2*g.Unit)
				}
			}
		}

		offset := 2
		for i := 0; i < g.Num_Thief+1; i++ {
			if g.Sprites[i+1].Motion.X < float64(-g.Unit/2+offset) {
				g.Sprites[i+1].Motion.X = float64(g.ScreenWidth + g.Unit/2 - offset)
			}
			if g.Sprites[i+1].Motion.X > float64(g.ScreenWidth+g.Unit/2-offset) {
				g.Sprites[i+1].Motion.X = float64(-g.Unit/2 + offset)
			}
		}

		for i := 0; i < g.Num_Sp; i++ {
			if g.Sprites[g.Num_Thief+i+2].Coll(g.Sprites[1]) {
				g.Sprites[g.Num_Thief+i+2].C = color.RGBA{100, 100, 100, 200}
			}
			for j := 0; j < g.Num_Thief; j++ {
				if g.Sprites[g.Num_Thief+i+2].Coll(g.Sprites[j+2]) {
					g.Sprites[g.Num_Thief+i+2].C = color.RGBA{255, 255, 0, 200}
				}
			}
		}

		fin_po, fin_th := true, true

		for i := 0; i < g.Num_Thief; i++ {
			if !(g.Sprites[i+2].Motion.X == center_x && g.Sprites[i+2].Motion.Y == center_y) {
				fin_po = false
			}
		}

		for i := 0; i < g.Num_Sp; i++ {
			r, _, _, _ := g.Sprites[g.Num_Thief+i+2].C.RGBA()
			if r == 25700 {
				fin_th = false
			}
		}

		if fin_po || fin_th {
			g.GameOver = true
		}
	} else {
		if ebiten.IsKeyPressed(ebiten.KeyQ) {
			po_x, po_y := 285, 225
			th_x, th_y := []float64{15, 555}, []float64{285, 285}

			g.Sprites[1].Motion.X = float64(po_x)
			g.Sprites[1].Motion.Y = float64(po_y)

			for i := 0; i < g.Num_Thief; i++ {
				th := g.Sprites[i+2]
				th.Motion.X = th_x[i]
				th.Motion.Y = th_y[i]
				th.Motion.MaxSpd = 2
			}

			for i := 0; i < g.Num_Sp; i++ {
				sp := g.Sprites[g.Num_Thief+i+2]
				sp.C = color.RGBA{100, 100, 100, 200}
			}
			g.GameOver = false
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	for _, i := range g.Sprites {
		if !i.Plable {
			i.Draw(screen)
		}
	}
	for i, _ := range g.Sprites {
		if g.Sprites[len(g.Sprites)-i-1].Plable {
			g.Sprites[len(g.Sprites)-i-1].Draw(screen)
		}
	}
}

func (g *Game) Layout(w, h int) (screenWidth, screenHeight int) {
	return g.ScreenWidth, g.ScreenHeight
}
