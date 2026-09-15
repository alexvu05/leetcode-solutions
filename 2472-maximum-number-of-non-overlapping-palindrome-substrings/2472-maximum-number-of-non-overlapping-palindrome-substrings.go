func maxPalindromes(s string, k int) int {
    n := len(s)
    count := 0
    i := 0  // leftmost available (not yet used) position
    for center := 0; center < 2*n-1; center++ {
        l := center / 2
        r := center / 2
        if center % 2 == 1 {
            r++  // even-length palindrome
        }
        // This center's leftmost character is l; skip if before available position
        if l < i { continue }
        // Expand from center outward
        for l >= 0 && r < n && s[l] == s[r] {
            curLen := r - l + 1
            if curLen >= k && l >= i {
                // Greedy: take this palindrome (shortest valid one from this center)
                count++
                i = r + 1
                break
            }
            l--
            r++
        }
    }
    return count
}