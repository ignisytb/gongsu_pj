package draw

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func pick_point(x, y int, c color.Color, screen *ebiten.Image) {
	screen.Set(x, y, c)
}
