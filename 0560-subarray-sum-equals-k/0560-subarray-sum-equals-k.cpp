class Solution {
public:
    int subarraySum(vector<int>& nums, int k) {
        // prefixCount[s] = number of times prefix sum s has been seen so far
        // Initialize with {0: 1} to handle subarrays starting from index 0
        unordered_map<int, int> prefixCount = {{0, 1}};
        int currentSum = 0;
        int count = 0;
        for (int x : nums) {
            currentSum += x;
            // Check how many previous prefix sums allow [j+1..i] to sum to k
            auto it = prefixCount.find(currentSum - k);
            if (it != prefixCount.end()) {
                count += it->second;
            }
            prefixCount[currentSum]++;
        }
        return count;
    }
};