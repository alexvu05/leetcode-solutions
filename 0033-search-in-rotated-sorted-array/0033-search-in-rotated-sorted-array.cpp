class Solution {
public:
    int search(vector<int>& nums, int target) {
        int left  = 0;
        int right = (int)nums.size() - 1;
        while (left <= right) {
            int mid = left + (right - left) / 2;  // avoid overflow vs (left+right)/2
            if (nums[mid] == target) return mid;
            // Determine which half is sorted
            if (nums[left] <= nums[mid]) {
                // Left half [left..mid] is sorted
                if (nums[left] <= target && target < nums[mid]) {
                    right = mid - 1;  // target lies in the sorted left half
                } else {
                    left = mid + 1;   // target must be in the right half
                }
            } else {
                // Right half [mid..right] is sorted
                if (nums[mid] < target && target <= nums[right]) {
                    left = mid + 1;   // target lies in the sorted right half
                } else {
                    right = mid - 1;  // target must be in the left half
                }
            }
        }
        return -1;  // target not found
    }
};