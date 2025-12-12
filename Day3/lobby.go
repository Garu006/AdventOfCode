// lobby
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Part selection: choose exactly k digits per line to maximize joltage.
	// Default k=12 for Part Two; you can override with -k.
	k := flag.Int("k", 12, "cantidad de dígitos a seleccionar por banco")
	flag.Parse()

	var input *os.File
	var err error

	// If a filepath is provided as the first argument, read from that file.
	// Otherwise, try to read from "input.txt" in the current folder.
	args := flag.Args()
	if len(args) > 0 {
		input, err = os.Open(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error abriendo archivo:", err)
			os.Exit(1)
		}
		defer input.Close()
	} else {
		input, err = os.Open("input.txt")
		if err != nil {
			// If opening input.txt fails, fall back to stdin
			input = os.Stdin
		} else {
			defer input.Close()
		}
	}

	Scanner := bufio.NewScanner(input)

	total := 0

	for Scanner.Scan() {
		line := strings.TrimSpace(Scanner.Text())
		if line == "" {
			continue // Saltar líneas vacías, por si acaso
		}

		bankMax := maxJoltageK(line, *k)
		total += bankMax
	}

	if err := Scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error de lectura:", err)
		os.Exit(1)
	}

	fmt.Println(total)

}

// maxJoltageK selects exactly k digits from line (preserving order)
// to form the largest possible number and returns its integer value.
func maxJoltageK(line string, k int) int {
	// Filter to digits only (input is digits per problem statement)
	digits := make([]byte, 0, len(line))
	for i := 0; i < len(line); i++ {
		c := line[i]
		if c >= '0' && c <= '9' {
			digits = append(digits, c)
		}
	}

	n := len(digits)
	if k <= 0 || n == 0 {
		return 0
	}
	if k >= n {
		// Use the whole number
		return parseIntBytes(digits)
	}

	// Remove r = n-k digits to maximize the number.
	r := n - k
	stack := make([]byte, 0, k)
	for i := 0; i < n; i++ {
		d := digits[i]
		// Greedy: while we can remove and current digit is larger
		for r > 0 && len(stack) > 0 && stack[len(stack)-1] < d {
			stack = stack[:len(stack)-1]
			r--
		}
		stack = append(stack, d)
	}
	// If still need to remove, remove from the end
	if r > 0 {
		stack = stack[:len(stack)-r]
	}
	// Only keep the first k digits (stack may be longer if no removals were needed at the end)
	if len(stack) > k {
		stack = stack[:k]
	}
	return parseIntBytes(stack)
}

func parseIntBytes(bs []byte) int {
	val := 0
	for i := 0; i < len(bs); i++ {
		val = val*10 + int(bs[i]-'0')
	}
	return val
}
