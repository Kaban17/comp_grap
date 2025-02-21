package parser

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Vertex struct {
	X, Y, Z float64
}

type Face struct {
	Indices []int
}

// ParseObj читает .obj файл и извлекает вершины и грани
func ParseObj(filename string) ([]Vertex, []Face, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var vertices []Vertex
	var faces []Face

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(line)

		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "v": // Вершина
			if len(parts) < 4 {
				continue
			}
			x, _ := strconv.ParseFloat(parts[1], 64)
			y, _ := strconv.ParseFloat(parts[2], 64)
			z, _ := strconv.ParseFloat(parts[3], 64)
			vertices = append(vertices, Vertex{X: x, Y: y, Z: z})

		case "f": // Грань (полигон)
			var indices []int
			for _, facePart := range parts[1:] {
				vertexIndex := strings.Split(facePart, "/")[0] // Берём только номер вершины
				idx, _ := strconv.Atoi(vertexIndex)
				indices = append(indices, idx)
			}
			faces = append(faces, Face{Indices: indices})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return vertices, faces, nil
}
