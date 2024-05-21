package game

import (
	"eng/sprite"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ScreenHeight int
	ScreenWidth  int
	Sprites      []*sprite.Sprites
}

func (g *Game) Update() error {
	Player := g.Sprites[0]
	if ebiten.IsKeyPressed(ebiten.KeyW) && Player.TouchFloor {
		if Player.JumpCount == 0 {
			Player.JumpCount = 13
			Player.Vl = 1000
			Player.Move(0, float64(Player.JumpCount))
			Player.Vl = 10
		} else if Player.JumpCount == 1 {
			Player.TouchFloor = false
			Player.JumpCount = 0
		} else {
			Player.JumpCount -= 1
			Player.Vl = 1000
			Player.Move(0, float64(Player.JumpCount))
			Player.Vl = 10
		}
	} else if Player.JumpCount != 0 {
		Player.JumpCount = 0
		Player.TouchFloor = false
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		Player.Move_Vel(-5, Player.Vy)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		Player.Move_Vel(5, Player.Vy)
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) && !Player.TouchFloor {
		Player.Vl = 1000
		Player.Move(0, -40)
		Player.Vl = 10
	}

	Player.Render()
	for _, sp := range g.Sprites {
		if sp.Gravity {
			sp.Move(0, -2)
		}
	}
	Player.Coll_wall(g.ScreenWidth, g.ScreenHeight)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	for _, sp := range g.Sprites {
		sp.Draw(screen)
	}
}

func (g *Game) Layout(W, H int) (screenWidth, screenHeight int) {
	return g.ScreenWidth, g.ScreenHeight
}
