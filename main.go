package main

import (
	"image"
	"image/color"
	"image/png"
	m "img/matrix"
	"img/parser"
	"log"
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
	vertices, faces, err := parser.ParseObj("model_1.obj")
	if err != nil {
		log.Fatal(err)
	}

	m := m.NewMatrix(2000, 2000, true, m.RGBColor{R: 255, G: 255, B: 255})

	m.DrawModel(vertices, faces)

	img := matrixToImage(&m)

	if err := saveImage(img, "attachement/output.png"); err != nil {
		log.Fatal(err)
	}

	log.Println("Изображение сохранено в attachement/output.png")
}
