package draw

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawRectangle(x, y, W, H int, c color.Color, screen *ebiten.Image) {
	for i := x; i < x+W+1; i++ {
		for j := y; j < y+H+1; j++ {
			pick_point(i, j, c, screen)
		}
	}
}

func DrawCircle(x, y, rad int, c color.Color, screen *ebiten.Image) {
	for i := x - rad; i < x+rad+1; i++ {
		for j := y - rad; j < y+rad+1; j++ {
			if (math.Pow(float64(i-x), 2) + math.Pow(float64(j-y), 2)) < math.Pow(float64(rad), 2) {
				pick_point(i, j, c, screen)
			}
		}
	}
}
