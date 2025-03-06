package triangle

import (
	"img/matrix"
	"math"
)

func (t *TriangleVertices) Draw(m *matrix.Matrix) {
	x_min := int(math.Min(t.X0, math.Min(t.X1, t.X2)))
	if x_min < 0 {
		x_min = 0
	}
	y_min := int(math.Min(t.Y0, math.Min(t.Y1, t.Y2)))
	if y_min < 0 {
		y_min = 0
	}
	x_max := int(math.Max(t.X0, math.Max(t.X1, t.X2)))
	if x_max < 0 {
		x_max = 0
	}
	y_max := int(math.Max(t.Y0, math.Max(t.Y1, t.Y2)))
	if y_max < 0 {
		y_max = 0
	}

	for y := y_min; y <= y_max; y++ {
		for x := x_min; x <= x_max; x++ {
			p := Point2D{X: float64(x), Y: float64(y)}
			b := p.Bar_coord(t)
			if b.X >= 0 && b.Y >= 0 && b.Z >= 0 {
				m.Set(x, y, matrix.RGBColor{R: 0, G: 0, B: 0})
			}
		}
	}
}
