package game

import (
	"eng/sprite"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ScreenHeight int
	ScreenWidth  int
	Unit         int
	Sprites      []*sprite.Sprites
}

func (g *Game) Update() error {
	player := g.Sprites[0]
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		player.Y -= player.CharSpd
		if player.Coll(g.Sprites[1]) {
			player.Y += player.CharSpd
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		player.X -= player.CharSpd
		if player.Coll(g.Sprites[1]) {
			player.X += player.CharSpd
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		player.Y += player.CharSpd
		if player.Coll(g.Sprites[1]) {
			player.Y -= player.CharSpd
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		player.X += player.CharSpd
		if player.Coll(g.Sprites[1]) {
			player.X -= player.CharSpd
		}
	}

	if player.X < (-player.Unit / 2) {
		player.X = g.ScreenWidth + (player.Unit / 2)
	}
	if player.X > g.ScreenWidth+(player.Unit/2) {
		player.X = (-player.Unit / 2)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	for _, sp := range g.Sprites {
		if !sp.Plable {
			sp.Draw(screen)
		}
	}
	for _, sp := range g.Sprites {
		if sp.Plable {
			sp.Draw(screen)
		}
	}
}

func (g *Game) Layout(W, H int) (screenWidth, screenHeight int) {
	return g.ScreenWidth, g.ScreenHeight
}
