func lexicographicallySmallestArray(nums []int, limit int) []int {
    n := len(nums)
    // Sort indices by value
    sortedIdx := make([]int, n)
    for i := range sortedIdx {
        sortedIdx[i] = i
    }
    sort.Slice(sortedIdx, func(i, j int) bool {
        return nums[sortedIdx[i]] < nums[sortedIdx[j]]
    })
    result := make([]int, n)
    // Process each group of indices that can mutually swap
    i := 0
    for i < n {
        j := i + 1
        // Extend group while consecutive sorted values differ by <= limit
        for j < n && nums[sortedIdx[j]]-nums[sortedIdx[j-1]] <= limit {
            j++
        }
        // Group = sortedIdx[i..j-1]
        // Collect their original indices, sort them
        groupIdx := make([]int, j-i)
        copy(groupIdx, sortedIdx[i:j])

        // Sort original indices to know which positions to fill
        sort.Ints(groupIdx)

        // Fill positions in sorted order with values in sorted order
        // sortedIdx[i..j-1] already has values in sorted order
        for k := 0; k < j-i; k++ {
            result[groupIdx[k]] = nums[sortedIdx[i+k]]
        }
        i = j
    }
    return result
}