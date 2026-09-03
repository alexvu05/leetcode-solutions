func uniformArray(nums1 []int) bool {
    // Find minimum odd number in nums1
    mn := math.MaxInt
    for _, x := range nums1 {
        if x%2 == 1 && x < mn {
            mn = x
        }
    }
    // If no odd number exists → all even → already uniform → true
    // If no even number exists → all odd → already uniform → true
    // If both exist: check every even x has at least one odd < x
    // Sufficient condition: min odd <= all even numbers
    // i.e., no even x < mn
    for _, x := range nums1 {
        if x%2 == 0 && mn != math.MaxInt && x < mn {
            return false
        }
    }
    return true
}