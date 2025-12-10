// lobby
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	Scanner := bufio.NewScanner(os.Stdin)

	total := 0

	for Scanner.Scan() {
		line := strings.TrimSpace(Scanner.Text())
		if line == "" {
			continue // Saltar líneas vacías, por si acaso
		}

		bankMax := maxJoltagbank(line)
		total += bankMax
	}

	if err := Scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error de lectura:", err)
		os.Exit(1)
	}

	fmt.Println(total)

}

func maxJoltagbank(line string) int {
	n := len(line)
	maxVal := -1

	for i := 0; i < n; i++ {
		c1 := line[i]
		if c1 < '0' || c1 > '9' {
			continue
		}
		d1 := int(c1 - '0')

		for j := i + 1; j < n; j++ {
			c2 := line[j]	
			if c2 < '0' || c2 > '9' {
				continue
			}
			d2 := int(c2 - '0')

			val := d1*10 + d2
			if val > maxVal {
				maxVal = val
			}
		}
	}

	if maxVal < 0 {
		return 0
	}
	return maxVal
}

