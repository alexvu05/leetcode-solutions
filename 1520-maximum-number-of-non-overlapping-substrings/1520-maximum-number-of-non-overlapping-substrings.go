func maxNumOfSubstrings(s string) []string {
    // Step 1: Find first and last occurrence of each char
    first := [26]int{}
    last  := [26]int{}
    for i := range first { first[i] = -1 }
    for i, c := range s {
        idx := c - 'a'
        if first[idx] == -1 { first[idx] = i }
        last[idx] = i
    }

    // Step 2: For each char, compute valid interval [l, r]
    // Expand: include all chars in [l,r] and their full ranges
    // Returns -1 if char not in s
    getInterval := func(c int) (int, int) {
        if first[c] == -1 { return -1, -1 }
        l, r := first[c], last[c]
        i := l
        for i <= r {
            ch := int(s[i] - 'a')
            if first[ch] < l { return -1, -1 } // char appears before l → invalid
            if last[ch] > r { r = last[ch] }   // expand r to include all of ch
            i++
        }
        return l, r
    }

    // Collect all valid intervals (one per starting char)
    type Interval struct{ l, r int }
    intervals := []Interval{}
    for c := 0; c < 26; c++ {
        l, r := getInterval(c)
        if l == -1 { continue }
        // Only add if this is a "root" interval (starts at first[c])
        if l == first[c] {
            intervals = append(intervals, Interval{l, r})
        }
    }

    // Step 3: Greedy — sort by right, pick non-overlapping
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].r < intervals[j].r
    })

    result := []string{}
    prevEnd := -1
    for _, iv := range intervals {
        if iv.l > prevEnd {
            result = append(result, s[iv.l:iv.r+1])
            prevEnd = iv.r
        }
    }

    return result
}