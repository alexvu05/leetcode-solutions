func countCommas(n int64) int64 {
    ans := int64(0)
    x := int64(1000)
    for x <= n {
        ans += n - x + 1
        x *= 1000
    }
    return ans
}