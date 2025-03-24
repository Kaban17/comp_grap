package triangle

import (
	"fmt"
	"img/matrix"
	"math"
)

func (t *TriangleVertices) Draw(m *matrix.Matrix) {
	x_min := int(math.Max(math.Min(t.X0, math.Min(t.X1, t.X2)), 0))
	y_min := int(math.Max(math.Min(t.Y0, math.Min(t.Y1, t.Y2)), 0))
	x_max := int(math.Max(t.X0, math.Max(t.X1, t.X2)))
	y_max := int(math.Max(t.Y0, math.Max(t.Y1, t.Y2)))

	for y := y_min; y <= y_max; y++ {
		for x := x_min; x <= x_max; x++ {
			p := Point2D{X: float64(x), Y: float64(y)}
			b, err := p.BarCoord(t)
			if err != nil {
				fmt.Println("Ошибка в BarCoord:", err)
				continue
			}

			if b.X >= 0 && b.Y >= 0 && b.Z >= 0 {
				if x >= 0 && y >= 0 && x < m.Cols && y < m.Rows {
					m.Set(x, y, matrix.RGBColor{R: 0, G: 0, B: 0})
				}
			}
		}
	}
}
