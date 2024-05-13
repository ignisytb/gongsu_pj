package game

import (
	"eng/draw"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ScreenHeight int
	ScreenWidth  int
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	draw.Shaker(g.ScreenWidth/2, g.ScreenHeight/2, 0, screen)
	draw.DrawCircle(g.ScreenWidth/2, g.ScreenHeight/2, 100, color.RGBA{0xff, 0, 0, 0xff}, screen)
}

func (g *Game) Layout(W, H int) (screenWidth, screenHeight int) {
	return g.ScreenWidth, g.ScreenHeight
}
