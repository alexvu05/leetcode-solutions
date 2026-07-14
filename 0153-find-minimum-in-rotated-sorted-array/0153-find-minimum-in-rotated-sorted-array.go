func findMin(nums []int) int {
    left, right := 0, len(nums)-1
    for left < right {
        mid := left + (right-left)/2  // avoid overflow
        if nums[mid] > nums[right] {
            // mid is in the larger left portion; minimum is in the right half
            left = mid + 1
        } else {
            // mid is in the smaller right portion; minimum is mid or to its left
            right = mid
        }
    }
    // left == right: converged to the minimum element
    return nums[left]
}