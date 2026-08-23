func sumGame(num string) bool {
    n := len(num)
    half := n / 2
    sumL, sumR := 0, 0
    cntL, cntR := 0, 0
    for i := 0; i < half; i++ {
        if num[i] == '?' {
            cntL++
        } else {
            sumL += int(num[i] - '0')
        }
    }
    for i := half; i < n; i++ {
        if num[i] == '?' {
            cntR++
        } else {
            sumR += int(num[i] - '0')
        }
    }
    // Bob wins iff 2*(sumL - sumR) == 9*(cntR - cntL)
    // Alice wins iff this is NOT satisfied
    return 2*(sumL-sumR) != 9*(cntR-cntL)
}