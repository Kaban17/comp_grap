package matrix

import (
	"fmt"
)

type Matrix struct {
	Rows, Cols int
	Data       [][]Color
}

func NewMatrix(rows, cols int, isRGB bool, value Color) Matrix {
	matrix := Matrix{
		Rows: rows,
		Cols: cols,
		Data: make([][]Color, rows),
	}
	for i := range matrix.Data {
		matrix.Data[i] = make([]Color, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if isRGB {
				matrix.Data[i][j] = value
			} else {
				// Заполняем матрицу оттенков серого
				matrix.Data[i][j] = value
			}
		}
	}
	return matrix
}

func (m *Matrix) IsRGB() bool {
	_, is_RGB := m.Data[0][0].(RGBColor)

	return is_RGB
}

// Вывод содержимого матрицы
func (m *Matrix) Print() {
	for _, row := range m.Data {
		for _, color := range row {
			fmt.Print(color.Get(), " ")
		}
		fmt.Println()
	}
}
func (m *Matrix) Get(row, col int) Color {
	if row < 0 || row >= m.Rows || col < 0 || col >= m.Cols {
		panic("Index out of range")
	}
	return m.Data[row][col]
}
func (m *Matrix) Set(row, col int, color Color) {
	if row < 0 || row >= m.Rows || col < 0 || col >= m.Cols {
		panic("Index out of range")
	}
	m.Data[row][col] = color
}

func (m *Matrix) Gradient() {
	rows, cols := m.Rows, m.Cols

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if m.IsRGB() {
				r := uint8(i * 255 / rows)
				g := uint8(j * 255 / cols)
				b := uint8(255)
				m.Data[i][j] = RGBColor{R: r, G: g, B: b}
			} else {
				grayValue := uint8(i * 255 / rows)
				m.Data[i][j] = GrayColor{Value: grayValue}
			}
		}
	}
}
