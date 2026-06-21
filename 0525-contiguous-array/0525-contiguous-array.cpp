class Solution {
public:
    int findMaxLength(vector<int>& nums) {
        // Map from prefix sum -> first index where this sum was seen
        // Initialize {0: -1} to handle subarrays starting from index 0
        unordered_map<int, int> firstSeen = {{0, -1}};
        int prefixSum = 0;
        int maxLen    = 0;
        for (int i = 0; i < (int)nums.size(); i++) {
            // Treat 0 as -1 and 1 as +1; equal count of 0s and 1s means sum = 0
            prefixSum += (nums[i] == 1) ? 1 : -1;
            auto it = firstSeen.find(prefixSum);
            if (it != firstSeen.end()) {
                // Same prefix sum seen before -> subarray between sums to 0
                maxLen = max(maxLen, i - it->second);
            } else {
                // Only record the first occurrence to maximize future subarray length
                firstSeen[prefixSum] = i;
            }
        }
        return maxLen;
    }
};