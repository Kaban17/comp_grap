package matrix

import (
	p "img/parser"
	"math"
)

func (m *Matrix) Bresenham(x0, y0, x1, y1 int) {
	dx := int(math.Abs(float64(x1 - x0)))
	dy := int(math.Abs(float64(y1 - y0)))
	sx := 1
	sy := 1

	if x0 > x1 {
		sx = -1
	}
	if y0 > y1 {
		sy = -1
	}

	err := dx - dy

	for {
		m.Set(x0, y0, RGBColor{R: 152, G: 118, B: 84})
		if x0 == x1 && y0 == y1 {
			break
		}

		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func (m *Matrix) DrawModel(vertices []p.Vertex, faces []p.Face) {
	scale := 5000.0
	offsetX, offsetY := 1000.0, 525.0

	// Преобразуем вершины в 2D-координаты
	transformed := make([][2]int, len(vertices))
	for i, v := range vertices {
		transformed[i] = [2]int{
			int(-scale*v.Y + offsetY),
			int(scale*v.X + offsetX),
		}
	}

	// Проходим по граням и соединяем вершины
	for _, face := range faces {
		n := len(face.Indices)
		for i := range n {
			v1 := transformed[face.Indices[i]-1]
			v2 := transformed[face.Indices[(i+1)%n]-1] // Замыкаем фигуры
			m.Bresenham(v1[0], v1[1], v2[0], v2[1])
		}
	}
}
