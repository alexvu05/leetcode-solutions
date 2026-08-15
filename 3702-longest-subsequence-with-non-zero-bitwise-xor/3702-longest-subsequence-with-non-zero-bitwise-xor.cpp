class Solution {
public:
    int longestSubsequence(vector<int>& nums) {
        int n = nums.size();
        int totalXor = 0;
        int cnt0 = 0; // count of zero elements
        for (int x : nums) {
            totalXor ^= x;
            if (x == 0) cnt0++;
        }
        // Case 1: XOR of all elements is non-zero → use all n elements
        if (totalXor != 0) return n;
        // Case 2: all elements are 0 → no valid subsequence
        if (cnt0 == n) return 0;
        // Case 3: XOR = 0 but has at least one non-zero element
        // Remove one non-zero element → XOR of remaining = that element ≠ 0
        return n - 1;
    }
};