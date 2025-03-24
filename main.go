package main

import (
	"fmt"
	m "img/matrix"
	s "img/save"
	tr "img/triangle"
)

func main() {
	m := m.NewMatrix(500, 500, true, m.RGBColor{R: 255, G: 255, B: 255})
	t := &tr.TriangleVertices{
		X0: 100.5, Y0: 100.0,
		X1: 203.1, Y1: 300.0,
		X2: 300.7, Y2: 100.01,
	}
	outpet_folder := "attachment"
	t.Draw(&m)
	img := s.MatrixToImage(&m)
	err := s.SaveImage(img, outpet_folder+"/triangle_1.png")
	if err != nil {
		fmt.Println("Error saving image:", err)
	}
	// vertex, faces, error := parser.ParseObj("model_1.obj")
	//
	//	if error != nil {
	//		fmt.Println("Error parsing obj:", error)
	//	}
	//
	//	for i := range 5 {
	//		fmt.Println(vertex[i], faces[i])
	//	}
}
