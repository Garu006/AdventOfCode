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

    // Part Two: iterative removal of accessible rolls.
    // Model as a graph of '@' nodes with edges to 8-neighbors. Repeatedly
    // remove nodes with degree < 4, updating neighbor degrees; count total removals.
    rows := len(grid)
    dirs := [][2]int{{-1,-1},{-1,0},{-1,1},{0,-1},{0,1},{1,-1},{1,0},{1,1}}

    // Degree matrix only for '@' cells
    deg := make([][]int, rows)
    for r := range deg {
        deg[r] = make([]int, cols)
    }

    // Compute initial degrees
    totalAt := 0
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] != '@' { continue }
            totalAt++
            n := 0
            for _, d := range dirs {
                nr, nc := r+d[0], c+d[1]
                if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '@' {
                    n++
                }
            }
            deg[r][c] = n
        }
    }

    // Queue of positions to remove (degree < 4)
    type pos struct{ r, c int }
    queue := make([]pos, 0)
    inQueue := make([][]bool, rows)
    removed := make([][]bool, rows)
    for r := 0; r < rows; r++ {
        inQueue[r] = make([]bool, cols)
        removed[r] = make([]bool, cols)
        for c := 0; c < cols; c++ {
            if grid[r][c] == '@' && deg[r][c] < 4 {
                queue = append(queue, pos{r, c})
                inQueue[r][c] = true
            }
        }
    }

    removedCount := 0
    // Process queue
    for head := 0; head < len(queue); head++ {
        p := queue[head]
        if removed[p.r][p.c] { // might be queued multiple times
            continue
        }
        // Remove this node
        removed[p.r][p.c] = true
        removedCount++
        // Decrement neighbors' degrees and enqueue newly accessible ones
        for _, d := range dirs {
            nr, nc := p.r+d[0], p.c+d[1]
            if nr < 0 || nr >= rows || nc < 0 || nc >= cols { continue }
            if grid[nr][nc] != '@' || removed[nr][nc] { continue }
            if deg[nr][nc] > 0 {
                deg[nr][nc]--
            }
            if deg[nr][nc] < 4 && !inQueue[nr][nc] {
                queue = append(queue, pos{nr, nc})
                inQueue[nr][nc] = true
            }
        }
    }

    // Output total removed
    fmt.Println(removedCount)
}