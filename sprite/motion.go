package sprite

type Motion struct {
	X, Y   float64
	Vx, Vy float64
	Ax, Ay float64
	MaxSpd float64
}

func (m *Motion) Update() {
	m.Friction()

	if m.Vx < m.MaxSpd && m.Vx > -m.MaxSpd {
		m.Vx += m.Ax
	}
	if m.Vy < m.MaxSpd && m.Vy > -m.MaxSpd {
		m.Vy += m.Ay
	}

	m.X += m.Vx
	m.Y += m.Vy
}

func (m *Motion) Friction() {
	friction := 0.1
	if m.Vx > 0 {
		m.Vx -= friction
	}
	if m.Vx < 0 {
		m.Vx += friction
	}
	if m.Vy > 0 {
		m.Vy -= friction
	}
	if m.Vy < 0 {
		m.Vy += friction
	}
}
