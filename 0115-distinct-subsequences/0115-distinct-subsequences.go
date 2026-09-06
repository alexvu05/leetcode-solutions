func numDistinct(s string, t string) int {
    m, n := len(s), len(t)
    // dp[j] = number of ways to match t[0..j-1] using s processed so far
    dp := make([]int, n+1)
    dp[0] = 1  // empty t always matches
    for i := 0; i < m; i++ {
        // Traverse right to left to avoid using s[i] twice
        for j := n; j >= 1; j-- {
            if s[i] == t[j-1] {
                dp[j] += dp[j-1]
            }
        }
    }
    return dp[n]
}