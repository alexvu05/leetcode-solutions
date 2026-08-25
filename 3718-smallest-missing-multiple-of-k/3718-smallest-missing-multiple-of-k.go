func missingMultiple(nums []int, k int) int {
    present := make(map[int]bool)
    for _, x := range nums {
        present[x] = true
    }
    // Check multiples of k starting from k
    for multiple := k; ; multiple += k {
        if !present[multiple] {
            return multiple
        }
    }
}