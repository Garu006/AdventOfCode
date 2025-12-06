// gift_shop.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// -------------------- main: lee la entrada y suma los IDs inválidos --------------------

func main() {
	// Leemos una sola línea de stdin (la del enunciado: a-b,c-d,...)
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	// En Advent of Code normalmente no hay salto de línea al final,
	// así que aceptamos EOF como algo normal.
	if err != nil && err.Error() != "EOF" {
		fmt.Fprintln(os.Stderr, "Error leyendo entrada:", err)
		return
	}

	line = strings.TrimSpace(line)
	if line == "" {
		fmt.Println(0)
		return
	}

	// Separamos los rangos por coma
	ranges := strings.Split(line, ",")

	var total int64 = 0

	for _, r := range ranges {
		r = strings.TrimSpace(r)
		if r == "" {
			continue
		}

		// Cada rango es "lo-hi"
		parts := strings.SplitN(r, "-", 2)
		if len(parts) != 2 {
			fmt.Fprintln(os.Stderr, "Rango inválido:", r)
			continue
		}

		lo, err1 := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
		hi, err2 := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
		if err1 != nil || err2 != nil {
			fmt.Fprintln(os.Stderr, "No pude convertir a int64:", r)
			continue
		}

		total += sumInvalidInRange(lo, hi)
	}

	fmt.Println(total)
}

// -------------------- utilidades numéricas --------------------

// 10^n como int64 (n pequeño, hasta ~18)
func pow10(n int) int64 {
	p := int64(1)
	for i := 0; i < n; i++ {
		p *= 10
	}
	return p
}

// Devuelve true si n tiene la forma “bloque repetido al menos 2 veces”.
// Ejemplos verdaderos: 55, 6464, 123123, 123123123, 1111111, 121212, etc.
func hasRepeatedPattern(n int64) bool {
	s := strconv.FormatInt(n, 10)
	if len(s) < 2 {
		return false
	}
	if s[0] == '0' {
		// No se usan IDs con cero inicial, por si acaso
		return false
	}
	L := len(s)

	// size = longitud del bloque base
	for size := 1; size*2 <= L; size++ { // al menos 2 bloques
		if L%size != 0 {
			continue
		}
		block := s[:size]
		ok := true
		for i := size; i < L; i += size {
			if s[i:i+size] != block {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}

// Suma todos los IDs inválidos en el rango [lo, hi] (Parte 2):
// números que son “bloque de dígitos repetido ≥ 2 veces”, sin duplicados.
func sumInvalidInRange(lo, hi int64) int64 {
	var total int64 = 0
	maxDigits := len(strconv.FormatInt(hi, 10))

	// h = número de dígitos del bloque base k
	for h := 1; h <= maxDigits; h++ {
		basePow := pow10(h)          // 10^h
		startK := pow10(h - 1)       // primer h-dígito (ej: 10, 100, 1000, ...)
		endK := basePow - 1          // último h-dígito  (ej: 99, 999, ...)
		for k := startK; k <= endK; k++ {

			// Si k ya es de por sí “repetición de algo más corto”,
			// lo saltamos para no generar duplicados.
			if hasRepeatedPattern(k) {
				continue
			}

			var rep int64 = 0

			// m = número de repeticiones del bloque k
			for m := 1; ; m++ {
				// Si el número total de dígitos excede maxDigits, ya no puede estar en el rango.
				if h*m > maxDigits {
					break
				}

				// Construimos el siguiente número repetido:
				// rep(m) = rep(m-1) * 10^h + k
				rep = rep*basePow + k

				// Solo desde m >= 2 es un ID inválido
				if m >= 2 {
					if rep > hi {
						// Si ya se pasó del rango, rep(m+1) será aún mayor → rompemos el bucle m
						break
					}
					if rep >= lo {
						total += rep
					}
				}
			}
		}
	}

	return total
}
