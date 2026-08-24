func stoneGameVIII(stones []int) int {
    n := len(stones)
    // Build prefix sums in-place
    for i := 1; i < n; i++ {
        stones[i] += stones[i-1]
    }
    // stones[i] now holds prefix sum s[i]
    // dp[n-1] = s[n-1] (Alice takes all stones, no choice for Bob)
    best := stones[n-1]
    // Iterate from n-2 down to 1 (Alice must take at least 2 stones → starts from index 1)
    for i := n - 2; i >= 1; i-- {
        // Current player can either:
        // 1. Take up to index i: score = s[i] - best (opponent gets best from i+1)
        // 2. Skip index i: score = best (let opponent choose from i+1)
        best = max(best, stones[i]-best)
    }
    return best
}