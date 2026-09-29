class Solution {
public:
    bool hasValidPath(vector<vector<char>>& grid) {
        int m = grid.size(), n = grid[0].size();
        if ((m + n - 1) % 2 != 0) return false;

        int maxBal = (m + n) / 2;  // max possible balance

        // dp[i][j] = bitset of reachable balances at (i,j)
        // Use vector<vector<bitset>> — but bitset size must be compile-time
        // Use uint32_t or just vector<bool> per cell
        // max balance <= 100 (m,n <= 100, so m+n <= 200, balance <= 100)
        const int MAXB = 101;
        vector<vector<vector<bool>>> dp(m, vector<vector<bool>>(n, vector<bool>(MAXB, false)));

        // Apply char: shift balance
        auto applyChar = [&](vector<bool>& src, vector<bool>& dst, char c) {
            fill(dst.begin(), dst.end(), false);
            if (c == '(') {
                for (int b = 0; b + 1 < MAXB; b++)
                    if (src[b]) dst[b+1] = true;
            } else {
                for (int b = 1; b < MAXB; b++)
                    if (src[b]) dst[b-1] = true;
            }
        };

        // Merge two bitsets (union)
        auto mergeBits = [&](vector<bool>& a, vector<bool>& b, vector<bool>& out) {
            for (int i = 0; i < MAXB; i++)
                out[i] = a[i] || b[i];
        };

        // Start
        vector<bool> start(MAXB, false);
        start[0] = true;
        applyChar(start, dp[0][0], grid[0][0]);

        // First row
        for (int j = 1; j < n; j++)
            applyChar(dp[0][j-1], dp[0][j], grid[0][j]);

        // First col
        for (int i = 1; i < m; i++)
            applyChar(dp[i-1][0], dp[i][0], grid[i][0]);

        // Rest
        vector<bool> tmp(MAXB);
        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) {
                mergeBits(dp[i-1][j], dp[i][j-1], tmp);
                applyChar(tmp, dp[i][j], grid[i][j]);
            }
        }

        return dp[m-1][n-1][0];
    }
};