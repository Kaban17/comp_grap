package triangle

import (
	"fmt"
	"img/matrix"
	"math"
)

type Point2D struct {
	X, Y float64
}
type TriangleVertices2D struct {
	P1, P2, P3 Point2D
	Color      matrix.RGBColor
}
type Point3D struct {
	X, Y, Z float64
}
type TriangleVertices3D struct {
	P1, P2, P3 Point3D
	Color      matrix.RGBColor
}

type ZBuffer [][]float64

func NewZBuffer(width, height int) ZBuffer {
	zb := make(ZBuffer, height)
	for y := range zb {
		zb[y] = make([]float64, width)
		for x := range zb[y] {
			zb[y][x] = math.MaxFloat64 // Инициализация большим значением
		}
	}
	return zb
}
func (p *Point2D) BarCoord(t *TriangleVertices2D) (Point3D, error) {
	denominator := (t.P1.X-t.P3.X)*(t.P2.Y-t.P3.Y) - (t.P2.X-t.P3.X)*(t.P1.Y-t.P3.Y)
	if denominator == 0 {
		return Point3D{}, fmt.Errorf("Треугольник вырожден")
	}

	lambda0 := ((p.X-t.P3.X)*(t.P2.Y-t.P3.Y) - (t.P2.X-t.P3.X)*(p.Y-t.P3.Y)) / denominator
	lambda1 := ((t.P1.X-t.P3.X)*(p.Y-t.P3.Y) - (p.X-t.P3.X)*(t.P1.Y-t.P3.Y)) / denominator
	lambda2 := 1.0 - lambda0 - lambda1

	return Point3D{X: lambda0, Y: lambda1, Z: lambda2}, nil
}
func (t *TriangleVertices3D) CalculateNormal() Point3D {
	v1 := Point3D{
		X: t.P2.X - t.P1.X,
		Y: t.P2.Y - t.P1.Y,
		Z: t.P2.Z - t.P1.Z,
	}
	v2 := Point3D{
		X: t.P3.X - t.P1.X,
		Y: t.P3.Y - t.P1.Y,
		Z: t.P3.Z - t.P1.Z,
	}

	nx := v1.Y*v2.Z - v1.Z*v2.Y
	ny := v1.Z*v2.X - v1.X*v2.Z
	nz := v1.X*v2.Y - v1.Y*v2.X

	length := math.Sqrt(nx*nx + ny*ny + nz*nz)
	if length > 0 {
		nx /= length
		ny /= length
		nz /= length
	}
	return Point3D{X: nx, Y: ny, Z: nz}
}
func (tri2D *TriangleVertices2D) DrawTriangleWithZBuffer(tri3D TriangleVertices3D, mat *matrix.Matrix, zb ZBuffer) {
	xMin := int(math.Max(math.Min(tri2D.P1.X, math.Min(tri2D.P2.X, tri2D.P3.X)), 0))
	yMin := int(math.Max(math.Min(tri2D.P1.Y, math.Min(tri2D.P2.Y, tri2D.P3.Y)), 0))
	xMax := int(math.Min(math.Max(tri2D.P1.X, math.Max(tri2D.P2.X, tri2D.P3.X)), float64(mat.Cols-1)))
	yMax := int(math.Min(math.Max(tri2D.P1.Y, math.Max(tri2D.P2.Y, tri2D.P3.Y)), float64(mat.Rows-1)))

	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			p := Point2D{X: float64(x) + 0.5, Y: float64(y) + 0.5}
			b, err := p.BarCoord(tri2D)
			if err != nil {
				continue
			}

			if b.X > 0 && b.Y > 0 && b.Z > 0 {
				// Вычисляем z-координату
				z := b.X*tri3D.P1.Z + b.Y*tri3D.P2.Z + b.Z*tri3D.P3.Z

				// Проверяем z-буфер
				if z < zb[y][x] {
					// Обновляем z-буфер и рисуем пиксель
					zb[y][x] = z
					mat.Set(x, y, tri2D.Color)
				}
			}
		}
	}
}
