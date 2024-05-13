package sprite

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprites struct {
	Img        *ebiten.Image
	Rot, Scale float64
	X, Y       float64
	Vx, Vy     float64
	Vl         float64
	Gravity    bool
	TouchFloor bool
	JumpCount  int
	Friction   float64
}

func (sp *Sprites) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	scale := sp.Scale
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(float64(sp.X)-float64(sp.Img.Bounds().Size().X/2)*scale, float64(sp.Y)-float64(sp.Img.Bounds().Size().Y/2)*scale)
	op.GeoM.Rotate(math.Pi * float64(sp.Rot) / 180)

	screen.DrawImage(sp.Img, op)
}

func (sp *Sprites) Move(Ax, Ay float64) {
	sp.Vx += Ax
	sp.Vy += Ay
	if sp.Vx > sp.Vl {
		sp.Vx = sp.Vl
	}
	if sp.Vx < -sp.Vl {
		sp.Vx = -sp.Vl
	}
	if sp.Vy > sp.Vl {
		sp.Vy = sp.Vl
	}
	if sp.Vy < -sp.Vl {
		sp.Vy = -sp.Vl
	}
}

func (sp *Sprites) Move_Vel(Vx, Vy float64) {
	sp.Vx = Vx
	sp.Vy = Vy
}

func (sp *Sprites) Render() {
	sp.X += sp.Vx
	sp.Y -= sp.Vy
	if sp.Vx > 0 {
		sp.Vx -= 0.4
	} else if sp.Vx < 0 {
		sp.Vx += 0.4
	}
	if sp.Vy > 0 {
		sp.Vy -= sp.Friction
	} else if sp.Vy < 0 {
		sp.Vy += sp.Friction
	}
}

func (sp *Sprites) Coll_wall(Width, Height int) {
	if float64(sp.X)-float64(sp.Img.Bounds().Size().X/2)*sp.Scale < 0 {
		sp.X = (float64(sp.Img.Bounds().Size().X/2) * sp.Scale)
	}
	if float64(sp.Y)-float64(sp.Img.Bounds().Size().Y/2)*sp.Scale < 0 {
		sp.Y = (float64(sp.Img.Bounds().Size().Y/2) * sp.Scale)
	}
	if float64(sp.X)+float64(sp.Img.Bounds().Size().X/2)*sp.Scale > float64(Width) {
		sp.X = float64(Width) - (float64(sp.Img.Bounds().Size().X/2) * sp.Scale)
	}
	if float64(sp.Y)+float64(sp.Img.Bounds().Size().Y/2)*sp.Scale > float64(Height) {
		sp.Y = float64(Height) - (float64(sp.Img.Bounds().Size().Y/2) * sp.Scale)
		sp.TouchFloor = true
	}
}

func Shaker(x, y, rot float64) *Sprites {
	img, _, err := ebitenutil.NewImageFromFile("sprite/shaker.png")
	if err != nil {
		fmt.Print("[DRAW/Shaker] shaker image loading error\n")
	}
	scale := 0.1

	return &Sprites{Img: img, X: x, Y: y, Vx: 0, Vy: 0, Vl: 10, Rot: rot, Scale: scale, Gravity: true, TouchFloor: false, Friction: 0.5}
}
