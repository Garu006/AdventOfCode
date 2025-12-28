package main

import (
	"bufio"	
	"fmt"
	"os"
	"strings"
)

func main() { // leer el archivo con readLineAndPad, Obtener los rangos con splitProblemsByColumns, y luego extraer y resolver cada subproblema con extractAndSolveSubProblems
	lines, maxlen, err := readLinesAndPad("input.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	// Obtener los rangos de columnas de cada problema
	problems := splitProblemsByColumns(lines, maxlen)
	// Extraer y resolver cada subproblema
	results := extractAndSolveSubProblems(lines, problems)
	fmt.Println("Results: ", results)

	// Sumar los resultados de cada subproblema
	total := 0
	for _, res := range results {
		total += res
	}
	fmt.Println("Total: ", total)
}

// readLinesAndPad read all lines, find maxLen,pad each line with spaces to maxLen, and return the lines and maxLen.

func readLinesAndPad(path string) ([]string, int, error) { // leer todas las lineas, encontrar maxLen, rellenar cada linea con espacios hasta maxLen, y devolver las lineas y maxLen.
	f, err := os.Open(path) // abre el archivo
	if err != nil { // si hay un error
		return nil, 0, err // devuelve nil, 0 y el error
	}
	defer f.Close() // cierra el archivo al final de la funcion

	scanner := bufio.NewScanner(f)	// crea un scanner para leer el archivo
	scanner.Buffer(make([]byte, 0, 1024) , 1024*1024) // aumentar el tamaño del buffer del scanner
	var lines []string // crea un slice para guardar las lineas
	maxLen := 0 // inicializa maxLen a 0
	for scanner.Scan() { // lee cada linea
		line := scanner.Text() // obtiene el texto de la linea
		lines = append(lines, line) // añade la linea al slice
		if len(line) > maxLen {// si la longitud de la linea es mayor que maxLen
			maxLen = len(line) // actualiza maxLen
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, 0, err
	}
	for i, line := range lines {
		if len(line) < maxLen {
			lines[i] = line + strings.Repeat(" ", maxLen - len(line))
		}
	}
	return lines, maxLen, nil
}

func isEmptyColumn(lines []string, col int) bool {
	for _, line := range lines {
		if line[col] != ' ' {
			return false
		}
	}
	return true
}

// lo que hace la funcion de abajo es devolver los rangos de columnas de cada problema

func splitProblemsByColumns(lines []string, maxLen int) [][2]int {
	var problems [][2]int
	inProblem := false
	var start int

	for col := 0; col < maxLen; col++ { // Recorremos cada columna
		if isEmptyColumn(lines, col) { // Si la columna esta vacia
			if inProblem {
				problems = append(problems, [2]int{start, col - 1}) // Añadimos el rango del problema
				inProblem = false // Marcamos que ya no estamsos en un problema
			}
		} else { // Si la columna no esta vaica
			if !inProblem {
				start = col // Marcamos el inicio del problema
				inProblem = true /// Marcamos que estamos en un problema
			}
		}
	}

	if inProblem { // Si al final seguimos en un problema
		problems = append(problems, [2]int{start, maxLen - 1}) // Añadimos el rango del problema
	}

	return problems // Devolvemos los rangos de problemas
}

func solveSubProblem(lines []string) int { // resolver el subproblema representado por lines
	var numbers []int // esta variable guarda los numeros encontrados
	op := "" // esta variable guarda la operacion encontrada

	for _, line := range lines { // recorremos cada linea
		line = strings.TrimSpace(line) // quitamos espacios en blanco al inicio y al final
		if line == "" { // si la linea esta vacia, seguimos con la siguiente
			continue 
		}

		if line == "+" || line == "*" { // si la linea tiene una operacion
			op = line // se guarda la operacion
		}else { // si la linea tiene un numero
			var n int // esta variable guarda el numero encontrado
			fmt.Sscanf(line, "%d", &n) // se convierte la linea a numero usando Sscanf
			numbers = append(numbers, n) // se agrega el numero a la lista de numeros
		} 
	}

	if len(numbers) == 0 || op == ""{ // si no se encontraron numeros o operaciones
		return 0 // se devuelve 0
	}

	result := numbers[0] // se inicializa el resultado con el primero numero
	for i := 1; i < len(numbers); i++ { // se recorren los demas numeros
		if op == "+" { // si la operacion es suma
			result += numbers[i] // se suma el numero al resultado
		}else { // si es multiplicacion
			result *= numbers[i] // se multiplica el numero al resultado
		}
	}
	return result // se devuelve el resultado
}

// Extraer cada subproblema usando los rangos de columnas y resolverlo de forma independiente

func extractAndSolveSubProblems(lines []string, problems [][2]int) []int {
	var results []int
	for _, p := range problems {
		sublines := make([]string, len(lines))
		for i, line := range lines {
			sublines[i] = line[p[0]:p[1]+1]
		}
		// Aquí va la lógica para resolver el subproblema representado por sublines
		result := solveSubProblem(sublines)
		results = append(results, result)
 	}
	return results
}
