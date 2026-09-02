func uniformArray(nums1 []int) bool {
    oddCnt, evenCnt := 0, 0
    for _, x := range nums1 {
        if x%2 == 0 {
            evenCnt++
        } else {
            oddCnt++
        }
    }
    // Can make all-odd?
    // - odd elements: keep as-is ✓
    // - even elements: need to subtract an odd → need oddCnt >= 1
    canAllOdd := evenCnt == 0 || oddCnt >= 1
    // Can make all-even?
    // - even elements: keep as-is ✓
    // - odd elements: need to subtract another odd (same parity diff = even)
    //   → need oddCnt >= 2 (since j != i)
    canAllEven := oddCnt == 0 || oddCnt >= 2
    return canAllOdd || canAllEven
}