class Solution {
public:
    int subarraysDivByK(vector<int>& nums, int k) {
        // remainderCount[r] = times prefix sum with remainder r has been seen
        unordered_map<int, int> remainderCount = {{0, 1}};
        int currentSum = 0;
        int count = 0;
        for (int x : nums) {
            currentSum += x;
            // Normalize remainder to handle negative values in C++
            int remainder = ((currentSum % k) + k) % k;
            // All previous prefix sums with same remainder form valid subarrays
            auto it = remainderCount.find(remainder);
            if (it != remainderCount.end()) {
                count += it->second;
            }
            remainderCount[remainder]++;
        }
        return count;
    }
};