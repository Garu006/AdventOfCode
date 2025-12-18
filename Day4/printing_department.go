package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 0, 64*1024), 10*1024*1024) // 10 MB max buffer
    var grid [][]byte
    for scanner.Scan() {
        line := scanner.Text()
        if line == "" { continue }
        grid = append(grid, []byte(line))
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "read error:", err)
        os.Exit(1)
    }

    if len(grid) == 0 {
        fmt.Println(0)
        return
    }
    cols := len(grid[0])
    for i := 1; i < len(grid); i++ {
        if len(grid[i]) != cols {
            fmt.Fprintln(os.Stderr, "Error: filas de distinto tamaño")
            os.Exit(1)
        }
    }

    dirs := [][2]int{{-1,-1},{-1,0},{-1,1},{0,-1},{0,1},{1,-1},{1,0},{1,1}}
    rows := len(grid) // reuse cols from above
    count := 0
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] != '@' { continue }
            n := 0
            for _, d := range dirs {
                nr, nc := r+d[0], c+d[1]
                if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '@' {
                    n++
                }
            }
            if n < 4 { count++ }
        }
    }
    fmt.Println(count)
}