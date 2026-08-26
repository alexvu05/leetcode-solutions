func shortestBeautifulSubstring(s string, k int) string {
    n := len(s)
    result := ""
    ones := 0  // count of '1' in current window
    left := 0
    for right := 0; right < n; right++ {
        if s[right] == '1' {
            ones++
        }
        // Shrink from left while we have k ones:
        // move left forward as long as s[left] == '0'
        // (removing '0' doesn't reduce ones count)
        for ones == k {
            // Try to shrink: remove leading zeros
            for left < right && s[left] == '0' {
                left++
            }
            // Current window s[left..right] has exactly k ones and no leading zero
            candidate := s[left : right+1]
            if result == "" ||
                len(candidate) < len(result) ||
                (len(candidate) == len(result) && candidate < result) {
                result = candidate
            }
            // Slide left past the leftmost '1' to try next window
            if s[left] == '1' {
                ones--
            }
            left++
        }
    }
    return result
}