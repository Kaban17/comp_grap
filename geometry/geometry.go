package geometry

import (
	"math"
)

type RGBColor struct {
	R, G, B uint8
}

type Matrix struct {
	Width, Height int
	Data          [][]RGBColor
	HasZBuffer    bool
}

func NewMatrix(width, height int, hasZBuffer bool, defaultColor RGBColor) Matrix {
	data := make([][]RGBColor, height)
	for i := range data {
		data[i] = make([]RGBColor, width)
		for j := range data[i] {
			data[i][j] = defaultColor
		}
	}
	return Matrix{
		Width:      width,
		Height:     height,
		Data:       data,
		HasZBuffer: hasZBuffer,
	}
}

type Point2D struct {
	X, Y float64
}

type Point3D struct {
	X, Y, Z float64
}

type TriangleVertices2D struct {
	P1, P2, P3 Point2D
	Color      RGBColor
}

type TriangleVertices3D struct {
	P1, P2, P3 Point3D
}

type Matrix3D [3][3]float64

func NewRotationMatrixY(angle float64) Matrix3D {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return Matrix3D{
		{cos, 0, sin},
		{0, 1, 0},
		{-sin, 0, cos},
	}
}

func NewRotationMatrixX(angle float64) Matrix3D {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return Matrix3D{
		{1, 0, 0},
		{0, cos, -sin},
		{0, sin, cos},
	}
}

func NewRotationMatrixZ(angle float64) Matrix3D {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return Matrix3D{
		{cos, -sin, 0},
		{sin, cos, 0},
		{0, 0, 1},
	}
}

func (m Matrix3D) Multiply(other Matrix3D) Matrix3D {
	var result Matrix3D
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				result[i][j] += m[i][k] * other[k][j]
			}
		}
	}
	return result
}

func (m Matrix3D) TransformPoint(p Point3D) Point3D {
	return Point3D{
		X: m[0][0]*p.X + m[0][1]*p.Y + m[0][2]*p.Z,
		Y: m[1][0]*p.X + m[1][1]*p.Y + m[1][2]*p.Z,
		Z: m[2][0]*p.X + m[2][1]*p.Y + m[2][2]*p.Z,
	}
}

func Translate(p Point3D, tx, ty, tz float64) Point3D {
	return Point3D{
		X: p.X + tx,
		Y: p.Y + ty,
		Z: p.Z + tz,
	}
}

func (t TriangleVertices3D) CalculateNormal() Point3D {
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

	normal := Point3D{
		X: v1.Y*v2.Z - v1.Z*v2.Y,
		Y: v1.Z*v2.X - v1.X*v2.Z,
		Z: v1.X*v2.Y - v1.Y*v2.X,
	}

	length := math.Sqrt(normal.X*normal.X + normal.Y*normal.Y + normal.Z*normal.Z)
	if length != 0 {
		normal.X /= length
		normal.Y /= length
		normal.Z /= length
	}

	return normal
}
