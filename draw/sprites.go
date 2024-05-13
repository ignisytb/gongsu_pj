package draw

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func Shaker(x, y int, rot float64, screen *ebiten.Image) {
	img, _, err := ebitenutil.NewImageFromFile("draw/shaker.png")
	if err != nil {
		fmt.Print("[DRAW/Shaker] shaker image loading error\n")
	}
	screen.DrawImage(img, nil)
}
