class Solution {
public:
    vector<int> findDuplicates(vector<int>& nums) {
        vector<int> duplicates;
        for (int i = 0; i < (int)nums.size(); i++) {
            // Each value v in [1..n] maps to index v-1.
            // We negate nums[v-1] on first visit.
            // If it's already negative, we've seen v before → duplicate.
            int idx = abs(nums[i]) - 1;
            if (nums[idx] < 0) {
                duplicates.push_back(idx + 1);
            } else {
                nums[idx] = -nums[idx];
            }
        }
        return duplicates;
    }
};