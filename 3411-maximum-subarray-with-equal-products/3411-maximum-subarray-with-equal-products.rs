impl Solution {
    // Time: O(n^2 * log C) | Space: O(1), C = max(nums) <= 10
    pub fn max_length(nums: Vec<i32>) -> i32 {
        fn gcd(a: i64, b: i64) -> i64 {
            if b == 0 { a } else { gcd(b, a % b) }
        }
        fn lcm(a: i64, b: i64) -> i64 {
            a / gcd(a, b) * b
        }

        let nums: Vec<i64> = nums.into_iter().map(|x| x as i64).collect();
        let n = nums.len();

        // gcd * lcm of any subarray never exceeds lcm(whole array) * max(nums)
        let max_val = *nums.iter().max().unwrap();
        let max_p = nums.iter().fold(1i64, |acc, &x| lcm(acc, x)) * max_val;

        let mut best = 0usize;
        for i in 0..n {
            let (mut p, mut g, mut l) = (1i64, 0i64, 1i64); // product, gcd (0 is identity), lcm
            for j in i..n {
                p *= nums[j];
                g = gcd(g, nums[j]);
                l = lcm(l, nums[j]);

                if p == g * l {
                    best = best.max(j - i + 1);
                }
                if p > max_p {
                    break; // product only grows → no longer extension can match
                }
            }
        }

        best as i32
    }
}