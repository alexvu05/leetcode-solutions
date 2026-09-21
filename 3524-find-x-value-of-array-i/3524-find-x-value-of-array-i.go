func resultArray(nums []int, k int) []int64 {
    result := make([]int64, k)

    // dp[v] = number of subarrays ending at current position with product % k == v
    dp := make([]int64, k)

    for _, x := range nums {
        xk := x % k
        newDp := make([]int64, k)

        // Extend existing subarrays: each subarray with product v → v*xk%k
        for v := 0; v < k; v++ {
            if dp[v] > 0 {
                newDp[(v*xk)%k] += dp[v]
            }
        }

        // Start new subarray at current position
        newDp[xk]++

        // Accumulate into result
        for v := 0; v < k; v++ {
            result[v] += newDp[v]
        }

        dp = newDp
    }

    return result
}