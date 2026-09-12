func maximumWeight(intervals [][]int) []int {
    n := len(intervals)
    // Attach original index before sorting
    type Interval struct{ l, r, w, idx int }
    ivs := make([]Interval, n)
    for i, iv := range intervals {
        ivs[i] = Interval{iv[0], iv[1], iv[2], i}
    }
    // Sort by right endpoint; ties: left asc, weight desc, original index asc
    sort.Slice(ivs, func(i, j int) bool {
        a, b := ivs[i], ivs[j]
        if a.r != b.r { return a.r < b.r }
        if a.l != b.l { return a.l < b.l }
        if a.w != b.w { return a.w > b.w }
        return a.idx < b.idx
    })
    rights := make([]int, n)
    for i, iv := range ivs {
        rights[i] = iv.r
    }
    const K = 4
    // dp[k][i] = {score, sorted indices} — best for k intervals using first i+1
    type State struct {
        score   int
        indices []int
    }
    // best[k][i] = best state for exactly k intervals, prefix [0..i]
    best := [K + 1][]State{}
    for k := 0; k <= K; k++ {
        best[k] = make([]State, n+1)
    }
    // Compare two states: prefer higher score; tie → lex smaller indices
    better := func(a, b State) bool {
        if a.score != b.score { return a.score > b.score }
        for i := 0; i < len(a.indices) && i < len(b.indices); i++ {
            if a.indices[i] != b.indices[i] {
                return a.indices[i] < b.indices[i]
            }
        }
        return false
    }
    for k := 1; k <= K; k++ {
        for i := 0; i < n; i++ {
            iv := ivs[i]
            // Don't take interval i: carry forward
            best[k][i+1] = best[k][i]
            // Binary search: last j where rights[j] < iv.l
            lo, hi, p := 0, i-1, -1
            for lo <= hi {
                mid := (lo + hi) / 2
                if rights[mid] < iv.l {
                    p = mid
                    lo = mid + 1
                } else {
                    hi = mid - 1
                }
            }
            // Take interval i: build new state
            prev := best[k-1][p+1]
            if k == 1 || prev.score > 0 || p >= 0 || k == 1 {
                newScore := prev.score + iv.w
                // Build new sorted indices
                newIndices := make([]int, len(prev.indices)+1)
                copy(newIndices, prev.indices)
                newIndices[len(prev.indices)] = iv.idx
                sort.Ints(newIndices)
                candidate := State{newScore, newIndices}
                // Also handle k=1 case where prev is empty
                if k == 1 && len(prev.indices) == 0 {
                    candidate = State{iv.w, []int{iv.idx}}
                }
                if best[k][i+1].score == 0 && len(best[k][i+1].indices) == 0 {
                    best[k][i+1] = candidate
                } else if better(candidate, best[k][i+1]) {
                    best[k][i+1] = candidate
                }
            }
        }
    }
    // Find best answer across all k=1..4
    var ans State
    for k := 1; k <= K; k++ {
        s := best[k][n]
        if len(s.indices) == 0 { continue }
        if len(ans.indices) == 0 || better(s, ans) {
            ans = s
        }
    }
    return ans.indices
}