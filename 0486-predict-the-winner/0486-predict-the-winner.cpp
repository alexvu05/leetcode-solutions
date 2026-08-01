class Solution {
public:
    bool predictTheWinner(vector<int>& nums) {
        int n = nums.size();
        // dp[i][j] = max score the current player can get from nums[i..j]
        vector<vector<int>> dp(n, vector<int>(n, 0));
        // Base case: single element
        for (int i = 0; i < n; i++) {
            dp[i][i] = nums[i];
        }
        // Fill by increasing interval length
        for (int len = 2; len <= n; len++) {
            for (int i = 0; i <= n - len; i++) {
                int j = i + len - 1;
                // Take nums[i]: opponent plays optimally on [i+1, j]
                int takeLeft  = nums[i] + min(
                    (i + 2 <= j ? dp[i+2][j] : 0),
                    (i + 1 <= j - 1 ? dp[i+1][j-1] : 0)
                );
                // Take nums[j]: opponent plays optimally on [i, j-1]
                int takeRight = nums[j] + min(
                    (i + 1 <= j - 1 ? dp[i+1][j-1] : 0),
                    (i <= j - 2 ? dp[i][j-2] : 0)
                );
                dp[i][j] = max(takeLeft, takeRight);
            }
        }
        int totalSum = accumulate(nums.begin(), nums.end(), 0);
        return 2 * dp[0][n-1] >= totalSum;
    }
};