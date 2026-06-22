class Solution {
public:
    int missingNumber(vector<int>& nums) {
        int n = nums.size();
        // Sum of 0..n using Gauss' formula
        long long expectedSum = (long long)n * (n + 1) / 2;
        long long actualSum = 0;
        for (int x : nums) {
            actualSum += x;
        }
        return (int)(expectedSum - actualSum);
    }
};