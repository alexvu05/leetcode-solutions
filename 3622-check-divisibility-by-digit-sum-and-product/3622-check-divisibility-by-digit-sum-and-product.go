func checkDivisibility(n int) bool {
    s, p := 0, 1  // digit sum, digit product
    x := n
    for x != 0 {
        v := x % 10   // extract last digit
        x /= 10       // remove last digit
        s += v
        p *= v
    }
    return n%(s+p) == 0
}