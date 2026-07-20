func shiftGrid(grid [][]int, k int) [][]int {
    m := len(grid)
    n := len(grid[0])
    total := m * n
    result := make([][]int, m)
    for i := range result {
        result[i] = make([]int, n)
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            // Current flat index
            pos := i*n + j
            // New flat index after k shifts
            newPos := (pos + k) % total
            // Convert back to 2D
            newRow := newPos / n
            newCol := newPos % n
            result[newRow][newCol] = grid[i][j]
        }
    }
    return result
}