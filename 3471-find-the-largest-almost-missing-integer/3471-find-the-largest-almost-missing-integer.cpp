class Solution {
public:
    int largestInteger(vector<int>& nums, int k) {
        int n = nums.size();
        if (k == 1) {
            unordered_map<int,int> freq;
            for (int x : nums) freq[x]++;
            int best = -1;
            for (auto& [x, c] : freq)
                if (c == 1) best = max(best, x);
            return best;
        }
        if (k == n) return *max_element(nums.begin(), nums.end());
        // 1 < k < n:
        // nums[0] is almost missing iff value nums[0] not in nums[1..n-1]
        // nums[n-1] is almost missing iff value nums[n-1] not in nums[0..n-2]
        unordered_set<int> allExceptFirst(nums.begin() + 1, nums.end());
        unordered_set<int> allExceptLast(nums.begin(), nums.end() - 1);
        int best = -1;
        if (!allExceptFirst.count(nums[0]))   best = max(best, nums[0]);
        if (!allExceptLast.count(nums[n-1]))  best = max(best, nums[n-1]);
        return best;
    }
};