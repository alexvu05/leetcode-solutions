class Solution {
public:
    vector<int> findMissingElements(vector<int>& nums) {
        // Load all existing numbers into a hash set for O(1) lookup
        unordered_set<int> present(nums.begin(), nums.end());
        int minVal = *min_element(nums.begin(), nums.end());
        int maxVal = *max_element(nums.begin(), nums.end());
        vector<int> result;
        // Scan full range; collect numbers absent from the set
        for (int x = minVal + 1; x < maxVal; x++) {
            if (present.find(x) == present.end()) {
                result.push_back(x);
            }
        }
        return result;  // already in sorted order since we scan left to right
    }
};