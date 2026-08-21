func findKthSmallest(coins []int, k int) int64 {
    // Step 1: Remove redundant coins (coin[i] redundant if any coin[j] | coin[i])
    filtered := []int{}
    for _, c := range coins {
        redundant := false
        for _, d := range coins {
            if d != c && c%d == 0 {
                redundant = true
                break
            }
        }
        if !redundant {
            filtered = append(filtered, c)
        }
    }
    n := len(filtered)
    const INF = int64(2e14)
    // Step 2: Precompute LCM for all non-empty subsets via bitmask
    subsetLCM := make([]int64, 1<<n)
    subsetLCM[0] = 1
    for mask := 1; mask < (1 << n); mask++ {
        lsb  := bits.TrailingZeros(uint(mask))  // index of lowest set bit
        prev := mask & (mask - 1)               // mask without lowest bit
        l    := lcm(subsetLCM[prev], int64(filtered[lsb]))
        if l > INF {
            l = INF + 1  // cap to avoid overflow
        }
        subsetLCM[mask] = l
    }
    // count(x) = number of distinct amounts <= x (inclusion-exclusion)
    count := func(x int64) int64 {
        total := int64(0)
        for mask := 1; mask < (1 << n); mask++ {
            lcmVal := subsetLCM[mask]
            if lcmVal > x {
                continue
            }
            popcount := bits.OnesCount(uint(mask))
            if popcount%2 == 1 {
                total += x / lcmVal  // add for odd-size subsets
            } else {
                total -= x / lcmVal  // subtract for even-size subsets
            }
        }
        return total
    }
    // Step 3: Binary search — smallest x s.t. count(x) >= k
    minCoin := int64(math.MaxInt64)
    for _, c := range filtered {
        if int64(c) < minCoin {
            minCoin = int64(c)
        }
    }
    left  := int64(1)
    right := int64(k) * minCoin
    for left < right {
        mid := left + (right-left)/2
        if count(mid) >= int64(k) {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}

func gcd(a, b int64) int64 {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func lcm(a, b int64) int64 {
    if a == 0 || b == 0 {
        return 0
    }
    return a / gcd(a, b) * b
}