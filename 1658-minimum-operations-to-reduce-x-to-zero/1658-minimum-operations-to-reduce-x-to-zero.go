func minOperations(nums []int, x int) int {
    // Time: O(n) | Space: O(1)
    total := 0
    for _, v := range nums { total += v }

    target := total - x
    if target < 0 { return -1 }
    if target == 0 { return len(nums) }

    n := len(nums)
    maxLen := -1
    curSum := 0
    left := 0

    for right := 0; right < n; right++ {
        curSum += nums[right]

        // Shrink from left while sum > target
        for curSum > target {
            curSum -= nums[left]
            left++
        }

        if curSum == target {
            if right-left+1 > maxLen {
                maxLen = right - left + 1
            }
        }
    }

    if maxLen == -1 { return -1 }
    return n - maxLen
}