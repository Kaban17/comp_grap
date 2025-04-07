package main

import (
	"fmt"
	"img/geometry"
	"img/parser"
	"img/save"
	"math"
)

func main() {
	const (
		aX = 2000.0
		aY = 2000.0
		u0 = 250.0
		v0 = 250.0
	)

	angleX := math.Pi / 6
	angleY := -math.Pi / 4
	angleZ := math.Pi / 3

	rotationX := geometry.NewRotationMatrixX(angleX)
	rotationY := geometry.NewRotationMatrixY(angleY)
	rotationZ := geometry.NewRotationMatrixZ(angleZ)

	rotation := rotationX.Multiply(rotationY.Multiply(rotationZ))

	tx, ty, tz := 0.0, 0.0, 10.0

	mat := geometry.NewMatrix(500, 500, true, geometry.RGBColor{R: 255, G: 255, B: 255})
	zb := geometry.NewZBuffer(500, 500)

	vertex, faces, err := parser.ParseObj("model.obj")
	if err != nil {
		fmt.Println("Error parsing obj:", err)
		return
	}

	for _, face := range faces {
		if len(face.Indices) < 3 {
			continue
		}
		vert0 := vertex[face.Indices[0]-1]
		vert1 := vertex[face.Indices[1]-1]
		vert2 := vertex[face.Indices[2]-1]

		p0 := geometry.Point3D{
			X: float64(vert0.X),
			Y: float64(vert0.Y),
			Z: float64(vert0.Z),
		}
		p1 := geometry.Point3D{
			X: float64(vert1.X),
			Y: float64(vert1.Y),
			Z: float64(vert1.Z),
		}
		p2 := geometry.Point3D{
			X: float64(vert2.X),
			Y: float64(vert2.Y),
			Z: float64(vert2.Z),
		}

		p0 = rotation.TransformPoint(p0)
		p1 = rotation.TransformPoint(p1)
		p2 = rotation.TransformPoint(p2)

		p0.X += tx
		p0.Y += ty
		p0.Z += tz
		p1.X += tx
		p1.Y += ty
		p1.Z += tz
		p2.X += tx
		p2.Y += ty
		p2.Z += tz

		tri3D := geometry.TriangleVertices3D{
			P1: p0,
			P2: p1,
			P3: p2,
		}

		normal := tri3D.CalculateNormal()
		cosTheta := normal.Z // Направление света [0,0,1]

		if cosTheta < 0 {

			tri2D := geometry.TriangleVertices2D{
				P1: geometry.Point2D{
					X: p0.X*aX + u0,
					Y: -p0.Y*aY + v0,
				},
				P2: geometry.Point2D{
					X: p1.X*aX + u0,
					Y: -p1.Y*aY + v0,
				},
				P3: geometry.Point2D{
					X: p2.X*aX + u0,
					Y: -p2.Y*aY + v0,
				},
				Color: geometry.RGBColor{
					R: uint8(-255 * cosTheta),
					G: 0,
					B: 0,
				},
			}

			tri2D.DrawTriangleWithZBuffer(tri3D, &mat, zb)
		}
	}

	img := save.MatrixToImage(&mat)
	if err := save.SaveImage(img, "output1.png"); err != nil {
		fmt.Println("Error saving image:", err)
	}
}
