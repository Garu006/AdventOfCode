package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	pos := 50   // posición inicial del dial
	count := 0 // cuántas veces pasa por 0

	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error al leer input.txt:", err)
		return
	}

	tokens := strings.Fields(string(data))

	for _, token := range tokens {
		direction := token[0]   // 'L' o 'R'
		stepsStr := token[1:]   // número

		var steps int
		_, err := fmt.Sscanf(stepsStr, "%d", &steps)
		if err != nil {
			fmt.Println("Error al convertir:", stepsStr)
			return
		}

		delta := 0
		if direction == 'R' {
			delta = 1
		} else if direction == 'L' {
			delta = -1
		} else {
			fmt.Println("Dirección inválida:", token)
			return
		}

		// Avanzar paso por paso
		for i := 0; i < steps; i++ {
			pos += delta
			pos = (pos + 100) % 100
			if pos == 0 {
				count++
			}
		}
	}

	fmt.Println(count)
}
