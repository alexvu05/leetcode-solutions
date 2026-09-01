func minMoves(classroom []string, energy int) int {
    m, n := len(classroom), len(classroom[0])
    // Index all litter cells
    litterIdx := map[[2]int]int{}
    numL := 0
    var startR, startC int
    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            switch classroom[r][c] {
            case 'S':
                startR, startC = r, c
            case 'L':
                litterIdx[[2]int{r, c}] = numL
                numL++
            }
        }
    }
    fullMask := (1 << numL) - 1
    // BFS state: (r, c, mask, energy)
    // Optimization: for each (r, c, mask), track max energy seen
    // Only enqueue if new energy > previously seen energy at this state
    type State struct {
        r, c, mask, eng int
    }
    // maxEnergy[r][c][mask] = max energy seen at this (pos, mask)
    maxEnergy := make([][][]int, m)
    for r := range maxEnergy {
        maxEnergy[r] = make([][]int, n)
        for c := range maxEnergy[r] {
            maxEnergy[r][c] = make([]int, 1<<numL)
            for i := range maxEnergy[r][c] {
                maxEnergy[r][c][i] = -1
            }
        }
    }
    queue := []State{{startR, startC, 0, energy}}
    maxEnergy[startR][startC][0] = energy
    steps := 0
    dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    for len(queue) > 0 {
        nextQueue := []State{}
        for _, s := range queue {
            if s.mask == fullMask {
                return steps
            }
            for _, d := range dirs {
                nr, nc := s.r+d[0], s.c+d[1]
                if nr < 0 || nr >= m || nc < 0 || nc >= n {
                    continue
                }
                if classroom[nr][nc] == 'X' {
                    continue
                }
                if s.eng == 0 {
                    continue // can't move
                }
                newEng := s.eng - 1
                newMask := s.mask
                cell := classroom[nr][nc]
                if cell == 'R' {
                    newEng = energy // reset to max
                } else if cell == 'L' {
                    if idx, ok := litterIdx[[2]int{nr, nc}]; ok {
                        newMask |= (1 << idx)
                    }
                }
                // Only enqueue if this state has better energy than before
                if newEng > maxEnergy[nr][nc][newMask] {
                    maxEnergy[nr][nc][newMask] = newEng
                    nextQueue = append(nextQueue, State{nr, nc, newMask, newEng})
                }
            }
        }
        queue = nextQueue
        steps++
    }
    return -1
}