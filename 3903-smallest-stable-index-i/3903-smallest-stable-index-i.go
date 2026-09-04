func firstStableIndex(nums []int, k int) int {
    n := len(nums)
    // Precompute suffix minimum
    suffMin := make([]int, n)
    suffMin[n-1] = nums[n-1]
    for i := n - 2; i >= 0; i-- {
        if nums[i] < suffMin[i+1] {
            suffMin[i] = nums[i]
        } else {
            suffMin[i] = suffMin[i+1]
        }
    }
    // Sweep left to right, tracking prefix maximum
    prefMax := 0
    for i, x := range nums {
        if x > prefMax {
            prefMax = x
        }
        if prefMax-suffMin[i] <= k {
            return i
        }
    }
    return -1
}