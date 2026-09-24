func smallestIndex(nums []int) int {
    for i, x := range nums {
        s := 0
        for x > 0 {
            s += x % 10
            x /= 10
        }
        if s == i {
            return i
        }
    }
    return -1
}