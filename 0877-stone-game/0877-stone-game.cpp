class Solution {
public:
    bool stoneGame(vector<int>& piles) {
        int n = piles.size();
        // dp[i] = max score current player can earn from current window starting at i
        vector<int> dp(piles.begin(), piles.end());
        for (int len = 2; len <= n; len++) {
            for (int i = 0; i <= n - len; i++) {
                int j = i + len - 1;
                // Take left or right, opponent gets the worse option
                dp[i] = max(piles[i] - dp[i+1],   // take left,  opponent has dp[i+1]
                            piles[j] - dp[i]);     // take right, opponent has dp[i]
            }
        }
        // dp[0] = Alice's score advantage over Bob
        return dp[0] > 0;
    }
};