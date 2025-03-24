package matrix

import "math"

func arange(start, stop, step float64) []float64 {
	if step == 0 {
		panic("step cannot be zero")
	}
	if (start < stop && step < 0) || (start > stop && step > 0) {
		return []float64{} // Возвращаем пустой массив, если шаг не соответствует направлению
	}

	var result []float64
	for value := start; (step > 0 && value < stop) || (step < 0 && value > stop); value += step {
		result = append(result, value)
	}
	return result
}

func (m *Matrix) Dotted_line(x0, y0, x1, y1, count int) {
	step := 1.0 / float64(count)

	for _, t := range arange(0.0, 1.0, step) {
		x := int(math.Round((1.0-t)*float64(x0) + t*float64(x1)))
		y := int(math.Round((1.0-t)*float64(y0) + t*float64(y1)))
		m.Set(x, y, RGBColor{R: 0, G: 0, B: 0}) // Устанавливаем точку на матрице
	}
}

func (m *Matrix) Dotted_linev2(x0, y0, x1, y1 int) {
	count := math.Sqrt(math.Pow(float64(x0-x1), 2) + math.Pow(float64(y0-y1), 2))
	step := 1.0 / float64(count)

	for _, t := range arange(0.0, 1.0, step) {
		x := int(math.Round((1.0-t)*float64(x0) + t*float64(x1)))
		y := int(math.Round((1.0-t)*float64(y0) + t*float64(y1)))
		m.Set(x, y, RGBColor{R: 0, G: 0, B: 0}) // Устанавливаем точку на матрице
	}
}

func (m *Matrix) LoopLines(x0, y0, x1, y1 int) {
	for x := x0; x < x1; x++ {
		t := float64(x-x0) / float64(x1-x0)
		y := int(math.Round((1.0-t)*float64(y0) + t*float64(y1)))
		m.Set(x, y, RGBColor{R: 0, G: 0, B: 0})
	}
}
func (m *Matrix) LoopLine_h1(x0, y0, x1, y1 int) {
	if x0 > x1 {
		x0, x1 = x1, x0
		y0, y1 = y1, y0
	}
	for x := x0; x < x1; x++ {
		t := float64(x-x0) / float64(x1-x0)
		y := int(math.Round((1.0-t)*float64(y0) + t*float64(y1)))
		m.Set(x, y, RGBColor{R: 0, G: 0, B: 0})
	}
}
func (m *Matrix) LoopLine_h2(x0, y0, x1, y1 int) {
	x_change := false
	if x0 > x1 {
		x0, x1 = x1, x0
		y0, y1 = y1, y0
	}
	if math.Abs(float64(x0-x1)) < math.Abs(float64(y0-y1)) {
		x0, y0 = y0, x0
		x1, y1 = y1, x1
		x_change = true
	}
	for x := x0; x <= x1; x++ {
		t := float64(x-x0) / float64(x1-x0)
		y := int(math.Round((1.0-t)*float64(y0) + t*float64(y1)))
		if x_change {
			m.Set(y, x, RGBColor{R: 0, G: 0, B: 0})
		} else {
			m.Set(x, y, RGBColor{R: 0, G: 0, B: 0})
		}
	}
}
