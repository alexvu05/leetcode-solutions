func minSumOfLengths(arr []int, target int) int {
    n := len(arr)
    const INF = 1 << 30
    // best[i] = min length of subarray with sum=target ending at or before index i
    best := make([]int, n)
    for i := range best { best[i] = INF }
    ans := INF
    curSum := 0
    left := 0
    for right := 0; right < n; right++ {
        curSum += arr[right]
        // Shrink window from left while sum > target
        for curSum > target {
            curSum -= arr[left]
            left++
        }
        if curSum == target {
            curLen := right - left + 1
            // Combine with best subarray ending strictly before 'left'
            if left > 0 && best[left-1] != INF {
                if curLen + best[left-1] < ans {
                    ans = curLen + best[left-1]
                }
            }
            // Update best[right]
            prev := INF
            if right > 0 { prev = best[right-1] }
            if curLen < prev {
                best[right] = curLen
            } else {
                best[right] = prev
            }
        } else {
            // No valid subarray ending at right
            if right > 0 {
                best[right] = best[right-1]
            }
        }
    }
    if ans == INF { return -1 }
    return ans
}