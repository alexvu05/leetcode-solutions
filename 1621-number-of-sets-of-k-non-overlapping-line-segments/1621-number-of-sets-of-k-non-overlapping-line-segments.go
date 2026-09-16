func numberOfSets(n int, k int) int {
    const MOD = 1_000_000_007

    total := n + k - 1  // choose 2k from total points
    r := 2 * k

    if r > total { return 0 }

    // Compute C(total, r) mod MOD using Pascal or modular inverse
    // Precompute factorials
    maxN := total + 1
    fact := make([]int, maxN)
    inv  := make([]int, maxN)
    fact[0] = 1
    for i := 1; i < maxN; i++ {
        fact[i] = fact[i-1] * i % MOD
    }

    // Modular inverse using Fermat's little theorem
    pow := func(base, exp, mod int) int {
        result := 1
        base %= mod
        for exp > 0 {
            if exp%2 == 1 { result = result * base % mod }
            base = base * base % mod
            exp /= 2
        }
        return result
    }

    inv[maxN-1] = pow(fact[maxN-1], MOD-2, MOD)
    for i := maxN - 2; i >= 0; i-- {
        inv[i] = inv[i+1] * (i+1) % MOD
    }

    comb := func(n, r int) int {
        if r < 0 || r > n { return 0 }
        return fact[n] * inv[r] % MOD * inv[n-r] % MOD
    }

    return comb(total, r)
}