func searchInsert(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        mid := left + (right-left)/2  // avoid overflow
        if nums[mid] == target {
            return mid           // exact match found
        } else if nums[mid] < target {
            left = mid + 1       // target is in the right half
        } else {
            right = mid - 1      // target is in the left half
        }
    }
    // left is the correct insert position when target is not found:
    // all elements to the left are < target, all to the right are > target
    return left
}