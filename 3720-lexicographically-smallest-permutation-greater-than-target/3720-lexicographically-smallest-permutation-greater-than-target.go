func lexGreaterPermutation(s string, target string) string {
    n := len(s)
    freq := [26]int{}
    for _, c := range s {
        freq[c-'a']++
    }
    // fillSorted fills result[pos..n-1] with remaining chars in sorted order
    fillSorted := func(result []byte, pos int, f [26]int) {
        idx := pos
        for ch := 0; ch < 26; ch++ {
            for f[ch] > 0 {
                result[idx] = byte('a' + ch)
                f[ch]--
                idx++
            }
        }
    }
    result := make([]byte, n)
    // Try each position i as the "branching point" where we place
    // a char STRICTLY GREATER than target[i], prefix [0..i-1] matches target exactly.
    // We go left to right: try to stay tight, at each position also try deviating.
    // To get lex smallest: prefer matching target[i] (stay tight) first,
    // then only deviate (place char > target[i]) if no better option downstream.
    //
    // Correct approach: iterate i from 0 to n-1.
    // At each i, maintaining tight prefix [0..i-1] = target[0..i-1]:
    //   Option A: place smallest char > target[i] at i, fill rest sorted → candidate
    //   Option B: place target[i] (stay tight), continue to i+1
    // We want lex smallest overall, so:
    //   If Option B leads to a valid answer, it's always <= Option A (same prefix, smaller/equal at i).
    //   So prefer Option B when possible (i.e., target[i] still available in freq).
    //   When Option B not possible (target[i] exhausted), must use Option A or fail.
    // After trying all tight positions, if none worked → return "".

    // Walk tight path, record each candidate from deviating at position i
    // but only return the LAST valid candidate (rightmost deviation = smallest result)
    // Actually: leftmost deviation gives largest prefix diff, rightmost gives smallest.
    // We want rightmost possible deviation point → smallest result.
    var best []byte
    // Simulate tight path
    tightFreq := freq
    for i := 0; i < n; i++ {
        t := int(target[i] - 'a')
        // Try to deviate at position i: place smallest char > target[i]
        for c := t + 1; c < 26; c++ {
            if tightFreq[c] > 0 {
                // Build candidate
                candidate := make([]byte, n)
                copy(candidate[:i], result[:i])
                candidate[i] = byte('a' + c)
                tmpFreq := tightFreq
                tmpFreq[c]--
                fillSorted(candidate, i+1, tmpFreq)
                // This is valid; record as best (later i → smaller result)
                best = candidate
                break
            }
        }
        // Stay tight: place target[i] if available
        if tightFreq[t] > 0 {
            result[i] = target[i]
            tightFreq[t]--
        } else {
            // Can't stay tight, no more options after this
            break
        }
    }
    return string(best)
}