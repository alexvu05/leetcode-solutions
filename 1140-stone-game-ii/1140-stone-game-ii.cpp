class Solution {
public:
    int stoneGameII(vector<int>& piles) {
        int n = piles.size();
        // Build suffix sums
        vector<int> suffix(n + 1, 0);
        for (int i = n - 1; i >= 0; i--) {
            suffix[i] = suffix[i + 1] + piles[i];
        }
        // dp[i][m] = max score current player can get from piles[i..n-1] with M=m
        // m is bounded by n since 2M grows fast and once 2M >= n it covers all
        vector<vector<int>> dp(n + 1, vector<int>(n + 1, 0));
        for (int i = n - 1; i >= 0; i--) {
            for (int m = 1; m <= n; m++) {
                // If can take all remaining piles
                if (i + 2 * m >= n) {
                    dp[i][m] = suffix[i];
                    continue;
                }
                // Try taking x = 1..2m piles
                for (int x = 1; x <= 2 * m; x++) {
                    int newM = max(m, x);
                    int score = suffix[i] - dp[i + x][newM];
                    dp[i][m] = max(dp[i][m], score);
                }
            }
        }
        return dp[0][1];
    }
};