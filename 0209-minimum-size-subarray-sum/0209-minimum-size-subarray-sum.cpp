class Solution {
public:
    int minSubArrayLen(int target, vector<int>& nums) {
        int left = 0;
        int currentSum = 0;
        int best = INT_MAX;
        for (int right = 0; right < (int)nums.size(); right++) {
            currentSum += nums[right];
            while (currentSum >= target) {
                best = min(best, right - left + 1);
                currentSum -= nums[left];
                left++;
            }
        }
        return best == INT_MAX ? 0 : best; 
    }
};