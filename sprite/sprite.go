package sprite

import (
	"eng/draw"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprites struct {
	X, Y, Unit int
	C          color.Color
	Plable     bool
	Hitmap     *MapAr
	CharSpd    int
}

func (s *Sprites) Draw(screen *ebiten.Image) {
	offset_x, offset_y := int(len(s.Hitmap.Mesh[0])*s.Unit/2), int(len(s.Hitmap.Mesh)*s.Unit/2)
	for i := range s.Hitmap.Mesh[0] {
		for j := range s.Hitmap.Mesh {
			if s.Hitmap.Mesh[j][i] {
				draw.DrawRectangle(i*s.Unit+s.X-offset_x, j*s.Unit+s.Y-offset_y, s.Unit, s.Unit, s.C, screen)
			}
		}
	}
}

func (s *Sprites) Coll_Point(x, y int) bool {
	i, j := x/s.Unit, y/s.Unit
	if i < 0 {
		i = 0
	}
	if i > (len(s.Hitmap.Mesh[0]) - 1) {
		i = len(s.Hitmap.Mesh[0]) - 1
	}
	if j < 0 {
		j = 0
	}
	if j > (len(s.Hitmap.Mesh) - 1) {
		j = len(s.Hitmap.Mesh) - 1
	}
	if s.Hitmap.Mesh[j][i] {
		return true
	} else {
		return false
	}
}

func (s *Sprites) Coll(oth *Sprites) bool {
	offset_x, offset_y := int(len(s.Hitmap.Mesh[0])*s.Unit/2), int(len(s.Hitmap.Mesh)*s.Unit/2)
	for i := range len(s.Hitmap.Mesh[0]) {
		for j := range len(s.Hitmap.Mesh) {
			for k := range s.Unit {
				for l := range s.Unit {
					if s.Hitmap.Mesh[j][i] {
						if oth.Coll_Point(i*s.Unit+k+s.X-offset_x, j*s.Unit+l+s.Y-offset_y) {
							return true
						}
					}
				}
			}
		}
	}
	return false
}
