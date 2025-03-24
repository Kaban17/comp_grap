package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	m "img/matrix"
	tr "img/triangle"
	"os"
)

func matrixToImage(matrix *m.Matrix) image.Image {
	rows, cols := matrix.Rows, matrix.Cols

	var img image.Image

	if matrix.IsRGB() {
		rgba := image.NewRGBA(image.Rect(0, 0, cols, rows))
		for y := 0; y < rows; y++ {
			for x := 0; x < cols; x++ {
				if rgb, ok := matrix.Data[y][x].(m.RGBColor); ok {
					rgba.Set(x, y, color.RGBA{R: rgb.R, G: rgb.G, B: rgb.B, A: 255})
				}
			}
		}
		img = rgba
	} else {
		gray := image.NewGray(image.Rect(0, 0, cols, rows))
		for y := 0; y < rows; y++ {
			for x := 0; x < cols; x++ {
				if grayPixel, ok := matrix.Data[y][x].(m.GrayColor); ok {
					gray.Set(x, y, color.Gray{Y: uint8(grayPixel.Value)})
				}
			}
		}
		img = gray
	}

	return img
}

func saveImage(img image.Image, filename string) error {
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return png.Encode(outFile, img)
}

func main() {
	m := m.NewMatrix(500, 500, true, m.RGBColor{R: 255, G: 255, B: 255})
	t := &tr.TriangleVertices{
		X0: 100.5, Y0: 100.0,
		X1: 203.1, Y1: 300.0,
		X2: -300.7, Y2: 100.01,
	}

	t.Draw(&m)
	img := matrixToImage(&m)
	err := saveImage(img, "attachment/triangle.png")
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
