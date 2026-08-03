class Solution {
public:
    string stoneGameIII(vector<int>& stoneValue) {
        int n = stoneValue.size();
        vector<int> dp(n + 3, 0);
        for (int i = n - 1; i >= 0; i--) {
            dp[i] = INT_MIN;
            int taken = 0;  // running sum of stones taken so far
            for (int take = 1; take <= 3 && i + take <= n; take++) {
                taken += stoneValue[i + take - 1];  // accumulate stones[i..i+take-1]
                int score = taken - dp[i + take];
                dp[i] = max(dp[i], score);
            }
        }
        if (dp[0] > 0) return "Alice";
        if (dp[0] < 0) return "Bob";
        return "Tie";
    }
};