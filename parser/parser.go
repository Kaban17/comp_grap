package parser

import (
	"bufio"
	"fmt"
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
		if len(line) == 0 {
			continue
		}

		// Разбор вершин
		if strings.HasPrefix(line, "v ") {
			parts := strings.Fields(line[2:])
			if len(parts) != 3 {
				return nil, nil, fmt.Errorf("неверный формат вершины: %s", line)
			}

			x, err := strconv.ParseFloat(parts[0], 64)
			if err != nil {
				return nil, nil, err
			}
			y, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return nil, nil, err
			}
			z, err := strconv.ParseFloat(parts[2], 64)
			if err != nil {
				return nil, nil, err
			}

			vertices = append(vertices, Vertex{x, y, z})
		}

		// Разбор граней
		if strings.HasPrefix(line, "f ") {
			parts := strings.Fields(line[2:])
			var indices []int
			for _, part := range parts {
				// Разбираем индексы вершин
				indicesParts := strings.Split(part, "/")
				index, err := strconv.Atoi(indicesParts[0])
				if err != nil {
					return nil, nil, err
				}
				indices = append(indices, index)
			}
			faces = append(faces, Face{Indices: indices})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return vertices, faces, nil
}
