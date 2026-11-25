impl Solution {
    pub fn smallest_unique_subarray(nums: Vec<i32>) -> i32 {
        const M: u64 = (1u64 << 61) - 1;

        fn mul(a: u64, b: u64) -> u64 {
            let t = a as u128 * b as u128;
            let mut r = ((t & M as u128) + (t >> 61)) as u64;
            if r >= M {
                r -= M;
            }
            r
        }

        let n = nums.len();

        // pseudo-random base from the clock
        let seed = std::time::SystemTime::now()
            .duration_since(std::time::UNIX_EPOCH)
            .map(|d| d.as_nanos() as u64)
            .unwrap_or(88172645463325252);
        let base = (seed % (M - 200_000)) + 100_001; // > max value (1e5)

        // prefix hashes and powers
        let mut h = vec![0u64; n + 1];
        let mut pw = vec![1u64; n + 1];
        for i in 0..n {
            h[i + 1] = (mul(h[i], base) + nums[i] as u64) % M;
            pw[i + 1] = mul(pw[i], base);
        }

        // does a subarray of length len appear exactly once?
        let has_unique = |len: usize| -> bool {
            let mut v: Vec<u64> = (0..=n - len)
                .map(|i| (h[i + len] + M - mul(h[i], pw[len])) % M)
                .collect();
            v.sort_unstable();
            let m = v.len();
            let mut i = 0;
            while i < m {
                let mut j = i + 1;
                while j < m && v[j] == v[i] {
                    j += 1;
                }
                if j - i == 1 {
                    return true;
                }
                i = j;
            }
            false
        };

        let (mut lo, mut hi) = (1usize, n);
        while lo < hi {
            let mid = (lo + hi) / 2;
            if has_unique(mid) {
                hi = mid;
            } else {
                lo = mid + 1;
            }
        }
        lo as i32
    }
}