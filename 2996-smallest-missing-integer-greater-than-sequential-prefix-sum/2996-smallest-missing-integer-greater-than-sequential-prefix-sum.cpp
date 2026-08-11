class Solution {
public:
    int missingInteger(vector<int>& nums) {
        // Step 1: compute sum of longest sequential prefix
        int sum = nums[0];
        for (int j = 1; j < (int)nums.size(); j++) {
            if (nums[j] == nums[j - 1] + 1) {
                sum += nums[j];
            } else {
                break;  // prefix ends here
            }
        }
        // Step 2: load all values into hash set
        unordered_set<int> present(nums.begin(), nums.end());
        // Step 3: find smallest x >= sum not in nums
        int x = sum;
        while (present.count(x)) {
            x++;
        }
        return x;
    }
};