package main

import (
	"fmt"
	"img/matrix"
	"img/parser"
	"img/save"
	tr "img/triangle"
)

func main() {
	const (
		scale   = 2500.0
		offsetX = 250.0
		offsetY = 250.0
	)

	// Создаем матрицу изображения и z-буфер
	mat := matrix.NewMatrix(500, 500, true, matrix.RGBColor{R: 255, G: 255, B: 255})
	zb := tr.NewZBuffer(500, 500)

	vertex, faces, err := parser.ParseObj("model.obj")
	if err != nil {
		fmt.Println("Error parsing obj:", err)
		return
	}

	for _, face := range faces {
		if len(face.Indices) < 3 {
			continue
		}

		v0 := vertex[face.Indices[0]-1]
		v1 := vertex[face.Indices[1]-1]
		v2 := vertex[face.Indices[2]-1]

		tri3D := tr.TriangleVertices3D{
			P1: tr.Point3D{X: v0.X, Y: v0.Y, Z: v0.Z},
			P2: tr.Point3D{X: v1.X, Y: v1.Y, Z: v1.Z},
			P3: tr.Point3D{X: v2.X, Y: v2.Y, Z: v2.Z},
		}

		normal := tri3D.CalculateNormal()
		cosTheta := normal.Z // Направление света [0,0,1]

		if cosTheta < 0 {
			tri2D := tr.TriangleVertices2D{
				P1: tr.Point2D{X: v0.X*scale + offsetX, Y: v0.Y*scale + offsetY},
				P2: tr.Point2D{X: v1.X*scale + offsetX, Y: v1.Y*scale + offsetY},
				P3: tr.Point2D{X: v2.X*scale + offsetX, Y: v2.Y*scale + offsetY},
				Color: matrix.RGBColor{
					R: uint8(-255 * cosTheta),
					G: 0,
					B: 0,
				},
			}

			tri2D.DrawTriangleWithZBuffer(tri3D, &mat, zb)
		}
	}

	img := save.MatrixToImage(&mat)
	if err := save.SaveImage(img, "output.png"); err != nil {
		fmt.Println("Error saving image:", err)
	}
}
