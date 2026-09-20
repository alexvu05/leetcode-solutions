func reverseDegree(s string) int {
    ans := 0
    for i, c := range s {
        reverseRank := 26 - int(c-'a')  // 'a'→26, 'z'→1
        ans += (i + 1) * reverseRank
    }
    return ans
}