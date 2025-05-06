package main

import (
	"fmt"
	"img/geometry"
	"img/parser"
	"img/save"
	"math"
)

func proj(p geometry.Point3D) geometry.Point2D {
	aX, aY := 5000.0, 5000.0
	u0, v0 := 250.0, 250.0
	return geometry.Point2D{
		X: aX*p.X/p.Z + u0,
		Y: aY*p.Y/p.Z + v0,
	}
}
func main() {
	const (
		angleX = math.Pi
		angleY = math.Pi / 4
		angleZ = 0
	)

	rotationX := geometry.NewRotationMatrixX(angleX)
	rotationY := geometry.NewRotationMatrixY(angleY)
	rotationZ := geometry.NewRotationMatrixZ(angleZ)

	rotation := rotationX.Multiply(rotationY.Multiply(rotationZ))

	tx, ty, tz := 0.0, 0.0, 2.0

	mat := geometry.NewMatrix(500, 500, true, geometry.RGBColor{R: 255, G: 255, B: 255})
	zb := geometry.NewZBuffer(500, 500)

	vertex, faces, err := parser.ParseObj("model.obj")
	if err != nil {
		fmt.Println("Error parsing obj:", err)
		return
	}
	var center geometry.Point3D
	for _, v := range vertex {
		center.X += v.X
		center.Y += v.Y
		center.Z += v.Z
	}
	n := float64(len(vertex))
	center.X /= n
	center.Y /= n
	center.Z /= n
	for _, face := range faces {
		if len(face.Indices) < 3 {
			continue
		}
		v0 := vertex[face.Indices[0]-1]
		v1 := vertex[face.Indices[1]-1]
		v2 := vertex[face.Indices[2]-1]

		p0 := geometry.Point3D{X: v0.X, Y: v0.Y, Z: v0.Z}
		p1 := geometry.Point3D{X: v1.X, Y: v1.Y, Z: v1.Z}
		p2 := geometry.Point3D{X: v2.X, Y: v2.Y, Z: v2.Z}

		p0 = rotation.TransformPoint(p0)
		p1 = rotation.TransformPoint(p1)
		p2 = rotation.TransformPoint(p2)

		p0 = geometry.Translate(p0, tx, ty, tz)
		p1 = geometry.Translate(p1, tx, ty, tz)
		p2 = geometry.Translate(p2, tx, ty, tz)

		p0 = geometry.Add(p0, center)
		p1 = geometry.Add(p1, center)
		p2 = geometry.Add(p2, center)
		tri3D := geometry.TriangleVertices3D{
			P1: geometry.Point3D{X: p0.X, Y: p0.Y, Z: p0.Z},
			P2: geometry.Point3D{X: p1.X, Y: p1.Y, Z: p1.Z},
			P3: geometry.Point3D{X: p2.X, Y: p2.Y, Z: p2.Z},
		}

		normal := tri3D.CalculateNormal()
		cosTheta := normal.Z

		if cosTheta < 0 {
			screenP1 := proj(p0)
			screenP2 := proj(p1)
			screenP3 := proj(p2)
			tri2D := geometry.TriangleVertices2D{
				P1: screenP1,
				P2: screenP2,
				P3: screenP3,
				Color: geometry.RGBColor{
					R: uint8(-255 * cosTheta),
					G: uint8(255 * cosTheta),
					B: uint8(255 * cosTheta),
				},
			}
			tri2D.DrawTriangleWithZBuffer(tri3D, &mat, zb)
		}
	}

	img := save.MatrixToImage(&mat)
	if err := save.SaveImage(img, "output2.png"); err != nil {
		fmt.Println("Error saving image:", err)
	}
}
