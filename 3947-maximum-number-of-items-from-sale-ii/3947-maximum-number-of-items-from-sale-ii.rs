impl Solution {
    // Time: O(n log n) | Space: O(n)
    pub fn maximum_sale_items(items: Vec<Vec<i32>>, budget: i32) -> i32 {
        let n = items.len();

        // cnt[v] = number of items whose factor is exactly v
        let mut cnt = vec![0i64; n + 1];
        for it in &items {
            cnt[it[0] as usize] += 1;
        }

        // mult[v] = number of items whose factor is a multiple of v (harmonic sieve)
        let mut mult = vec![0i64; n + 1];
        for v in 1..=n {
            let mut m = v;
            while m <= n {
                mult[v] += cnt[m];
                m += v;
            }
        }

        let mn = items.iter().map(|it| it[1]).min().unwrap() as i64; // cheapest price

        // (price, c): c = number of OTHER items unlocked as free by buying this item.
        // Only "double" purchases cheaper than 2*mn can beat two copies of the cheapest item.
        let mut cand: Vec<(i64, i64)> = items
            .iter()
            .map(|it| (it[1] as i64, mult[it[0] as usize] - 1))
            .filter(|&(p, c)| c > 0 && p < 2 * mn)
            .collect();
        cand.sort();                                   // cheapest first

        let mut rem = budget as i64;
        let mut copies = 0i64;
        for (p, c) in cand {
            let q = c.min(rem / p);                    // how many double-value copies we can afford
            copies += 2 * q;                           // each = 1 bought + 1 free
            rem -= q * p;
            if q < c {
                break;                                 // leftover < p ≤ every later price
            }
        }

        (copies + rem / mn) as i32                     // spend the rest on the cheapest item
    }
}