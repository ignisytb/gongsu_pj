package sprite

import (
	"eng/draw"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprites struct {
	Unit          int
	X, Y, CharSpd float64
	C             color.Color
	Plable        bool
	Hitmap        *MapAr
	Jail_Ent      time.Time
}

func (s *Sprites) Police_Move(wall *Sprites) bool {
	move := false
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		s.Y -= s.CharSpd
		if s.Coll(wall) {
			s.Y += s.CharSpd
		} else {
			move = true
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		s.Y += s.CharSpd
		if s.Coll(wall) {
			s.Y -= s.CharSpd
		} else {
			move = true
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		s.X -= s.CharSpd
		if s.Coll(wall) {
			s.X += s.CharSpd
		} else {
			move = true
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		s.X += s.CharSpd
		if s.Coll(wall) {
			s.X -= s.CharSpd
		} else {
			move = true
		}
	}
	return move
}

func (s *Sprites) Thief_Move(wall *Sprites) bool {
	move := false
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		s.Y -= s.CharSpd
		if s.Coll(wall) {
			s.Y += s.CharSpd
		} else {
			move = true
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		s.Y += s.CharSpd
		if s.Coll(wall) {
			s.Y -= s.CharSpd
		} else {
			move = true
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		s.X -= s.CharSpd
		if s.Coll(wall) {
			s.X += s.CharSpd
		} else {
			move = true
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		s.X += s.CharSpd
		if s.Coll(wall) {
			s.X -= s.CharSpd
		} else {
			move = true
		}
	}
	return move
}

func (s *Sprites) Draw(screen *ebiten.Image) {
	offset_x, offset_y := float64(len(s.Hitmap.Mesh[0])*s.Unit/2), float64(len(s.Hitmap.Mesh)*s.Unit/2)
	for i := 0; i < len(s.Hitmap.Mesh[0]); i++ {
		for j := 0; j < len(s.Hitmap.Mesh); j++ {
			if s.Hitmap.Mesh[j][i] {
				draw.DrawRectangle(int(s.X-offset_x)+i*s.Unit, int(s.Y-offset_y)+j*s.Unit, s.Unit, s.Unit, s.C, screen)
			}
		}
	}
}

func (s *Sprites) Coll_Point(x, y float64) bool {
	offset_x, offset_y := float64(len(s.Hitmap.Mesh[0])*s.Unit/2), float64(len(s.Hitmap.Mesh)*s.Unit/2)
	x_s, y_s := s.X-offset_x, s.Y-offset_y

	i, j := (x-x_s)/float64(s.Unit), (y-y_s)/float64(s.Unit)

	if i <= 0 || i >= float64(len(s.Hitmap.Mesh[0])) || j <= 0 || j >= float64(len(s.Hitmap.Mesh)) {
		return false
	}

	return s.Hitmap.Mesh[int(j)][int(i)]
}

func (s *Sprites) Coll(oth *Sprites) bool {
	offset_x, offset_y := float64(len(s.Hitmap.Mesh[0])*s.Unit/2), float64(len(s.Hitmap.Mesh)*s.Unit/2)
	x_s, y_s := s.X-offset_x, s.Y-offset_y

	for i := 0; i < len(s.Hitmap.Mesh[0]); i++ {
		for j := 0; j < len(s.Hitmap.Mesh); j++ {
			if s.Hitmap.Mesh[j][i] {
				for a := 0; a < s.Unit; a++ {
					for b := 0; b < s.Unit; b++ {
						if oth.Coll_Point(float64(i*s.Unit+a)+x_s, float64(j*s.Unit+b)+y_s) {
							return true
						}
					}
				}
			}
		}
	}
	return false
}
