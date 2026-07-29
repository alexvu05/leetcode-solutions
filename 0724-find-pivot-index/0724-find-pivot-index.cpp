class Solution {
public:
    int pivotIndex(vector<int>& nums) {
        int totalSum = accumulate(nums.begin(), nums.end(), 0);
        int leftSum  = 0;
        for (int i = 0; i < (int)nums.size(); i++) {
            // rightSum = totalSum - leftSum - nums[i]
            // pivot condition: leftSum == rightSum
            if (leftSum == totalSum - leftSum - nums[i]) {
                return i;
            }
            leftSum += nums[i];
        }
        return -1;
    }
};