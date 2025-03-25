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
