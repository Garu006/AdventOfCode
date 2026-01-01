package main

import (
	"bufio"
	"fmt"
	"os"
)

type resultadoRayo struct {
	filaFinal int
	colFinal int
	razonFin string
	salioMapa bool
}

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
	
	// iniciar simulacion de rayo desde la posicion S
	SimularRayo(startRow+1, startCol, lines)

	// imprimir resultados de la simulacion
	resultado := SimularRayo(startRow+1, startCol, lines)
	fmt.Printf("Rayo termina en fila %d, columna %d, razon: %s, salioMapa: %t\n", resultado.filaFinal, resultado.colFinal, resultado.razonFin, resultado.salioMapa)
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

// "devolver resultados" = entregar datos, no imprimir texto, el rayo debe decir: donde termino, por que termino y si salio del mapa
func SimularRayo(row, col int, lines []string) resultadoRayo {
	height := len(lines)

	// Dirección inicial: hacia abajo
	dFila, dCol := 1, 0

	for {
		// verificar si el rayo salió del mapa
		if row < 0 || row >= height || col < 0 || col >= len(lines[row]) {
			return resultadoRayo{
				filaFinal: row,
				colFinal:  col,
				razonFin:  "Salio del mapa",
				salioMapa: true,
			}
		}

		celda := lines[row][col]

		switch celda {
		case ' ', 'S':
			// continuar recto
			row += dFila
			col += dCol

		case '/':
			// reflejo /
			dFila, dCol = -dCol, -dFila
			row += dFila
			col += dCol

		case '\\':
			// reflejo \
			dFila, dCol = dCol, dFila
			row += dFila
			col += dCol

		case 'E':
			return resultadoRayo{
				filaFinal: row,
				colFinal:  col,
				razonFin:  "Llego a la salida",
				salioMapa: false,
			}

		default:
			return resultadoRayo{
				filaFinal: row,
				colFinal:  col,
				razonFin:  "Choco con un obstaculo",
				salioMapa: false,
			}
		}
	}
}
