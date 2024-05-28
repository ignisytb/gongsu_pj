package sprite

import (
	"bufio"
	"image/color"
	"os"
)

type MapAr struct {
	Mesh [][]bool
}

func (m *MapAr) Load_file(filename string) {
	data, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer data.Close()

	reader := bufio.NewReader(data)
	for {
		line, isPrefix, err := reader.ReadLine()
		if isPrefix || err != nil {
			break
		}
		nM := make([]bool, len(line))
		for i := range line {
			if line[i] == 49 {
				nM[i] = true
			}
			if line[i] == 48 {
				nM[i] = false
			}
		}
		m.Mesh = append(m.Mesh, [][]bool{nM}...)
	}
}

func (m *MapAr) Size() (w, h int) {
	return len(m.Mesh[0]), len(m.Mesh)
}

func Player(s int) *Sprites {
	ret := &MapAr{}

	ret.Mesh = [][]bool{{true}}

	return &Sprites{X: 0, Y: 0, Unit: s, C: color.CMYK{100, 200, 30, 30}, Hitmap: ret, Plable: true, CharSpd: 2}
}
