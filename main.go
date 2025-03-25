package main

import (
	"fmt"
	"img/matrix"
	"img/parser"
	s "img/save"
	tr "img/triangle"
	"math"
	"math/rand"
	"time"
)

func RandomColor() matrix.RGBColor {
	return matrix.RGBColor{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const (
		scale   = 2500.0
		offsetX = 250.0
		offsetY = 250.0
	)

	m := matrix.NewMatrix(500, 500, true, matrix.RGBColor{R: 255, G: 255, B: 255})

	output_folder := "attachment"

	vertex, faces, error := parser.ParseObj("model.obj")
	if error != nil {
		fmt.Println("Error parsing obj:", error)
		return
	}

	for i, face := range faces {
		if len(face.Indices) >= 3 {
			valid := true
			for _, idx := range face.Indices {
				adjustedIdx := idx - 1
				if adjustedIdx < 0 || adjustedIdx >= len(vertex) {
					fmt.Printf("Invalid vertex index %d in face %d\n", idx, i)
					valid = false
					break
				}
			}

			if valid {
				v0 := vertex[face.Indices[0]-1]
				v1 := vertex[face.Indices[1]-1]
				v2 := vertex[face.Indices[2]-1]

				// Создаем 3D-треугольник для расчета нормали
				tri3D := tr.TriangleVertices3D{
					P1: tr.Point3D{X: v0.X, Y: v0.Y, Z: v0.Z},
					P2: tr.Point3D{X: v1.X, Y: v1.Y, Z: v1.Z},
					P3: tr.Point3D{X: v2.X, Y: v2.Y, Z: v2.Z},
				}

				// Рассчитываем нормаль и косинус угла
				normal := tri3D.CalculateNormal()
				cosTheta := normal.Z // Т.к. свет направлен вдоль Z

				// Отрисовываем только полигоны с отрицательным косинусом
				if cosTheta < 0 {
					// Вычисляем интенсивность красного канала
					redIntensity := -255 * cosTheta // cosTheta < 0 => redIntensity > 0
					color := matrix.RGBColor{
						R: uint8(math.Round(redIntensity)),
						G: 0,
						B: 0,
					}

					// Создаем 2D-треугольник с рассчитанным цветом
					t := &tr.TriangleVertices2D{
						P1: tr.Point2D{
							X: v0.X*scale + offsetX,
							Y: v0.Y*scale + offsetY,
						},
						P2: tr.Point2D{
							X: v1.X*scale + offsetX,
							Y: v1.Y*scale + offsetY,
						},
						P3: tr.Point2D{
							X: v2.X*scale + offsetX,
							Y: v2.Y*scale + offsetY,
						},
						Color: color,
					}
					t.Draw(&m)
				}
			}
		}
	}

	img := s.MatrixToImage(&m)
	err := s.SaveImage(img, output_folder+"/model_2.png")
	if err != nil {
		fmt.Println("Error saving image:", err)
	}
}
