func lexPalindromicPermutation(s string, target string) string {
    n := len(s)
    half := n / 2

    // Step 1: Count freq, check palindrome possible
    freq := [26]int{}
    for _, c := range s {
        freq[c-'a']++
    }

    oddCount := 0
    midChar := byte(0)
    halfFreq := [26]int{}
    for i := 0; i < 26; i++ {
        halfFreq[i] = freq[i] / 2
        if freq[i]%2 == 1 {
            oddCount++
            midChar = byte('a' + i)
        }
    }

    // Palindrome impossible if more than 1 odd-frequency char
    if oddCount > 1 {
        return ""
    }

    // Build palindrome from half-string
    buildPalin := func(h []byte) string {
        res := make([]byte, n)
        copy(res[:half], h)
        if n%2 == 1 {
            res[half] = midChar
        }
        for i := 0; i < half; i++ {
            res[n-1-i] = h[i]
        }
        return string(res)
    }

    // fillSorted fills res[pos..] with remaining chars in f, sorted ascending
    fillSorted := func(res []byte, pos int, f [26]int) {
        idx := pos
        for ch := 0; ch < 26; ch++ {
            for f[ch] > 0 {
                res[idx] = byte('a' + ch)
                f[ch]--
                idx++
            }
        }
    }

    // Step 2: Greedy on half — same logic as bài 3720
    targetHalf := target[:half]
    halfResult := make([]byte, half)
    var bestHalf []byte
    tightFreq := halfFreq

    for i := 0; i < half; i++ {
        t := int(targetHalf[i] - 'a')

        // Try deviate at position i: place smallest char > target[i]
        for c := t + 1; c < 26; c++ {
            if tightFreq[c] > 0 {
                candidate := make([]byte, half)
                copy(candidate[:i], halfResult[:i])
                candidate[i] = byte('a' + c)
                tmpFreq := tightFreq
                tmpFreq[c]--
                fillSorted(candidate, i+1, tmpFreq)
                bestHalf = candidate
                break
            }
        }

        // Stay tight
        if tightFreq[t] > 0 {
            halfResult[i] = targetHalf[i]
            tightFreq[t]--
        } else {
            break
        }
    }

    // Step 3: Check if tight path completed (halfResult == targetHalf)
    // Count remaining in tightFreq
    remaining := 0
    for _, v := range tightFreq {
        remaining += v
    }
    if remaining == 0 {
        // Tight path = targetHalf → palindrome = buildPalin(halfResult)
        palin := buildPalin(halfResult)
        if palin > target {
            if bestHalf == nil || palin < buildPalin(bestHalf) {
                bestHalf = halfResult
            }
        }
    }

    if bestHalf == nil {
        return ""
    }

    palin := buildPalin(bestHalf)
    if palin > target {
        return palin
    }
    return ""
}