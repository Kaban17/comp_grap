package save

import (
	"image"
	"image/color"
	"image/png"
	m "img/matrix"
	"os"
)

func MatrixToImage(matrix *m.Matrix) image.Image {
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

func SaveImage(img image.Image, filename string) error {
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return png.Encode(outFile, img)
}
