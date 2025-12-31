package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines, width, err := readLines("input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	height := len(lines)
	fmt.Println("height: ", height, "width: ", width)

	startRow, startCol := -1, -1

	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if lines[r][c] == 'S' {
				startRow = r
				startCol = c
			}
		}
	}

	fmt.Println("S en: ", startRow, startCol)

	// startCol y starRow no deben ser -1, la posicion debe coincidir con el input
	if startRow == -1 || startCol == -1 {
		fmt.Println("Error: start position not found")
		return
	}
	
	SimularRayo(startRow+1, startCol, lines)
}

// leer el input y guardarlo commo una matriz accesible por fila y columna
func readLines(path string) ([]string, int, error) { // leer todas las lineas, encontrar maxLen, rellenar cada linea con espacios hasta maxLen, y devolver las lineas y maxLen.
	f, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	width := 0

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		if width == 0 {
			width = len(line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, 0, err
	}
	return lines, width, nil
}

func SimularRayo(row, col int, lines []string) {
	height := len(lines)

	for row < height {
		char := lines[row][col]
		if char == '.' {
			row++
			continue
		} else if char == '^' {
			fmt.Println("Rayo detenido en: ", row, col)
			return
		} else {
			fmt.Println("Rayo detenido por simbolo: ", string(char), "en", row, col)
			return
		}
	}
}