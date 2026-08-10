class Solution {
public:
    bool winnerSquareGame(int n) {
        // dp[i] = true if current player wins with i stones remaining
        vector<bool> dp(n + 1, false);
        // dp[0] = false: no stones left → current player loses
        for (int i = 1; i <= n; i++) {
            // Try removing each perfect square k² <= i
            for (int k = 1; k * k <= i; k++) {
                // If removing k² puts opponent in a losing position → we win
                if (!dp[i - k * k]) {
                    dp[i] = true;
                    break;  // found a winning move, no need to check further
                }
            }
        }
        return dp[n];
    }
};