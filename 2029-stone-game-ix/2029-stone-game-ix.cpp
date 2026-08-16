class Solution {
public:
    bool stoneGameIX(vector<int>& stones) {
        int cnt[3] = {};
        for (int s : stones) cnt[s % 3]++;
        if (cnt[0] % 2 == 0) {
            // Even zeros: Alice wins iff both type-1 and type-2 exist
            return cnt[1] > 0 && cnt[2] > 0;
        } else {
            // Odd zeros: Alice wins iff |cnt[1] - cnt[2]| > 2
            return abs(cnt[1] - cnt[2]) > 2;
        }
    }
};