package geometry

import "math"

type ZBuffer struct {
	Width, Height int
	Data          [][]float64
}

func NewZBuffer(width, height int) ZBuffer {
	data := make([][]float64, height)
	for i := range data {
		data[i] = make([]float64, width)
		for j := range data[i] {
			data[i][j] = 1e9
		}
	}
	return ZBuffer{
		Width:  width,
		Height: height,
		Data:   data,
	}
}

func (t TriangleVertices2D) DrawTriangleWithZBuffer(tri3D TriangleVertices3D, mat *Matrix, zb ZBuffer) {

	minX := int(math.Min(math.Min(t.P1.X, t.P2.X), t.P3.X))
	maxX := int(math.Max(math.Max(t.P1.X, t.P2.X), t.P3.X))
	minY := int(math.Min(math.Min(t.P1.Y, t.P2.Y), t.P3.Y))
	maxY := int(math.Max(math.Max(t.P1.Y, t.P2.Y), t.P3.Y))

	minX = max(0, minX)
	maxX = min(mat.Width-1, maxX)
	minY = max(0, minY)
	maxY = min(mat.Height-1, maxY)

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {

			w1 := ((t.P2.Y-t.P3.Y)*float64(x) + (t.P3.X-t.P2.X)*float64(y) + t.P2.X*t.P3.Y - t.P3.X*t.P2.Y) /
				((t.P2.Y-t.P3.Y)*t.P1.X + (t.P3.X-t.P2.X)*t.P1.Y + t.P2.X*t.P3.Y - t.P3.X*t.P2.Y)
			w2 := ((t.P3.Y-t.P1.Y)*float64(x) + (t.P1.X-t.P3.X)*float64(y) + t.P3.X*t.P1.Y - t.P1.X*t.P3.Y) /
				((t.P3.Y-t.P1.Y)*t.P2.X + (t.P1.X-t.P3.X)*t.P2.Y + t.P3.X*t.P1.Y - t.P1.X*t.P3.Y)
			w3 := 1 - w1 - w2

			if w1 >= 0 && w2 >= 0 && w3 >= 0 {

				z := w1*tri3D.P1.Z + w2*tri3D.P2.Z + w3*tri3D.P3.Z

				if z <= zb.Data[y][x] {
					zb.Data[y][x] = z
					mat.Data[y][x] = t.Color
				}
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
