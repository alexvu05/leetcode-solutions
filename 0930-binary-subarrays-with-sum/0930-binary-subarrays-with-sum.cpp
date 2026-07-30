class Solution {
public:
    int numSubarraysWithSum(vector<int>& nums, int goal) {
        // prefixCount[s] = number of times prefix sum s has been seen
        unordered_map<int, int> prefixCount = {{0, 1}};
        int currentSum = 0;
        int count = 0;
        for (int x : nums) {
            currentSum += x;
            // Check how many previous prefix sums allow [j+1..i] to sum to goal
            auto it = prefixCount.find(currentSum - goal);
            if (it != prefixCount.end()) {
                count += it->second;
            }
            prefixCount[currentSum]++;
        }
        return count;
    }
};