func largestOverlap(img1 [][]int, img2 [][]int) int {
    n := len(img1)
    // Collect positions of 1s in each image
    ones1 := [][2]int{}
    ones2 := [][2]int{}
    for r := 0; r < n; r++ {
        for c := 0; c < n; c++ {
            if img1[r][c] == 1 { ones1 = append(ones1, [2]int{r, c}) }
            if img2[r][c] == 1 { ones2 = append(ones2, [2]int{r, c}) }
        }
    }
    // Count frequency of each translation offset
    offsetCount := map[[2]int]int{}
    for _, p1 := range ones1 {
        for _, p2 := range ones2 {
            offset := [2]int{p2[0] - p1[0], p2[1] - p1[1]}
            offsetCount[offset]++
        }
    }
    best := 0
    for _, cnt := range offsetCount {
        if cnt > best { best = cnt }
    }
    return best
}