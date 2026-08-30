func minimumDeletions(nums []int) int {
    n := len(nums)
    // Find indices of min and max
    minIdx, maxIdx := 0, 0
    for i := 1; i < n; i++ {
        if nums[i] < nums[minIdx] {
            minIdx = i
        }
        if nums[i] > nums[maxIdx] {
            maxIdx = i
        }
    }
    // Strategy 1: remove both from left
    leftOnly := max(minIdx, maxIdx) + 1
    // Strategy 2: remove both from right
    rightOnly := max(n-1-minIdx, n-1-maxIdx) + 1
    // Strategy 3: remove one from each side
    // leftIdx = closer to left, rightIdx = closer to right
    leftIdx  := min(minIdx, maxIdx)
    rightIdx := max(minIdx, maxIdx)
    mixed := (leftIdx + 1) + (n - rightIdx)
    return min(leftOnly, min(rightOnly, mixed))
}