class Solution {
public:
    int stoneGameV(vector<int>& stoneValue) {
        int n = stoneValue.size();
        // Build prefix sums for O(1) range sum queries
        vector<int> prefix(n + 1, 0);
        for (int i = 0; i < n; i++) {
            prefix[i + 1] = prefix[i] + stoneValue[i];
        }
        auto rangeSum = [&](int l, int r) {
            return prefix[r + 1] - prefix[l];
        };
        // dp[i][j] = max score Alice can earn from stoneValue[i..j]
        vector<vector<int>> dp(n, vector<int>(n, 0));
        // Fill by increasing interval length
        for (int len = 2; len <= n; len++) {
            for (int i = 0; i <= n - len; i++) {
                int j = i + len - 1;
                for (int p = i; p < j; p++) {
                    int leftSum  = rangeSum(i, p);
                    int rightSum = rangeSum(p + 1, j);
                    if (leftSum < rightSum) {
                        // Bob throws away right; Alice scores leftSum + recurse on left
                        dp[i][j] = max(dp[i][j], leftSum + dp[i][p]);
                    } else if (leftSum > rightSum) {
                        // Bob throws away left; Alice scores rightSum + recurse on right
                        dp[i][j] = max(dp[i][j], rightSum + dp[p + 1][j]);
                    } else {
                        // Equal: Alice picks the better sub-problem
                        dp[i][j] = max(dp[i][j],
                            leftSum + max(dp[i][p], dp[p + 1][j]));
                    }
                }
            }
        }
        return dp[0][n - 1];
    }
};