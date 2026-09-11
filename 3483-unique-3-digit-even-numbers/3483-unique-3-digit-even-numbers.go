func totalNumbers(digits []int) int {
    // Count frequency of each digit in input
    freq := [10]int{}
    for _, d := range digits {
        freq[d]++
    }
    count := 0
    // Enumerate all 3-digit even numbers
    for x := 100; x <= 998; x += 2 {
        d1 := x / 100      // hundreds digit (must be 1-9)
        d2 := x / 10 % 10  // tens digit
        d3 := x % 10       // units digit (must be even)
        // Check if we can form this number from digits
        // Use a temporary freq copy to handle duplicate digits
        tmp := freq
        if tmp[d1] == 0 {
            continue
        }
        tmp[d1]--
        if tmp[d2] == 0 {
            continue
        }
        tmp[d2]--
        if tmp[d3] == 0 {
            continue
        }
        count++
    }
    return count
}