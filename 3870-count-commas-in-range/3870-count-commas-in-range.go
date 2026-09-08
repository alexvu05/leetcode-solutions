func countCommas(n int) int {
    ans := 0
    x := 1000  // 10^3
    for x <= n {
        // Numbers in [x, n] each have at least one more comma than [x/1000, x-1]
        ans += n - x + 1
        x *= 1000
    }
    return ans
}