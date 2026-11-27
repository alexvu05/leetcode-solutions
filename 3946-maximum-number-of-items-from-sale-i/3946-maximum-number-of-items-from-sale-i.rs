impl Solution {
    // Time: O(n^2 + n * budget) | Space: O(budget)
    pub fn maximum_sale_items(items: Vec<Vec<i32>>, budget: i32) -> i32 {
        let budget = budget as usize;
        let n = items.len();
        let mn = items.iter().map(|it| it[1]).min().unwrap() as usize; // cheapest price

        // dp[b] = max copies from "first purchases" (distinct types) with total cost <= b
        let mut dp = vec![0i32; budget + 1];

        for i in 0..n {
            let factor = items[i][0];
            let price = items[i][1] as usize;

            // free copies unlocked by buying item i: other items whose factor is a multiple of it
            let frees = items
                .iter()
                .enumerate()
                .filter(|(j, it)| *j != i && it[0] % factor == 0)
                .count() as i32;
            let value = 1 + frees;                       // the purchased copy + the free ones

            // 0/1 knapsack: iterate budget downwards so item i is used at most once
            if price <= budget {
                for b in (price..=budget).rev() {
                    dp[b] = dp[b].max(dp[b - price] + value);
                }
            }
        }

        // spend `i` on first purchases, put everything left into the cheapest item
        (0..=budget)
            .map(|i| dp[i] + ((budget - i) / mn) as i32)
            .max()
            .unwrap()
    }
}