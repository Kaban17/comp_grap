package save

import (
	"image"
	"image/color"
	"image/png"
	"img/geometry"
	"os"
)

func MatrixToImage(matrix *geometry.Matrix) image.Image {
	rows, cols := matrix.Height, matrix.Width

	rgba := image.NewRGBA(image.Rect(0, 0, cols, rows))
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			rgb := matrix.Data[y][x]
			rgba.Set(x, y, color.RGBA{R: rgb.R, G: rgb.G, B: rgb.B, A: 255})
		}
	}
	return rgba
}

func SaveImage(img image.Image, filename string) error {
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return png.Encode(outFile, img)
}
