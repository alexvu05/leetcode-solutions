func distinctSubseqII(s string) int {
    const MOD = 1_000_000_007
    // last[c] = dp value at the previous occurrence of char c
    // (the number of distinct subsequences BEFORE that occurrence was processed)
    last := [26]int{}  // initialized to 0
    dp := 0  // number of distinct non-empty subsequences so far
    for _, c := range s {
        idx := c - 'a'
        // New dp = 2*dp + 1 - last[idx]
        // Explanation:
        //   2*dp + 1 : append c to all existing subseqs, plus c alone
        //   - last[idx]: remove duplicates introduced by previous occurrence of c
        newDp := (2*dp + 1 - last[idx] + MOD) % MOD
        // Save current dp before updating, for future occurrences of c
        last[idx] = (dp + 1) % MOD  // dp+1 because we include c-alone subseq
        dp = newDp
    }
    return dp
}