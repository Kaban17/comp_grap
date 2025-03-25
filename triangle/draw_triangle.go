package triangle

import (
	"img/matrix"
	"math"
)

func (t *TriangleVertices2D) Draw(m *matrix.Matrix) {
	x_min := int(math.Max(math.Min(t.P1.X, math.Min(t.P2.X, t.P3.X)), 0))
	y_min := int(math.Max(math.Min(t.P1.Y, math.Min(t.P2.Y, t.P3.Y)), 0))
	x_max := int(math.Max(t.P1.X, math.Max(t.P2.X, t.P3.X)))
	y_max := int(math.Max(t.P1.Y, math.Max(t.P2.Y, t.P3.Y)))

	for y := y_min; y <= y_max; y++ {
		for x := x_min; x <= x_max; x++ {
			p := Point2D{X: float64(x), Y: float64(y)}
			b, err := p.BarCoord(t)
			if err != nil {
				continue
			}

			if b.X >= 0 && b.Y >= 0 && b.Z >= 0 {
				if x >= 0 && y >= 0 && x < m.Cols && y < m.Rows {
					m.Set(x, y, t.Color)
				}
			}
		}
	}
}
