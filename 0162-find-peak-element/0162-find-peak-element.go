func findPeakElement(nums []int) int {
    left, right := 0, len(nums)-1
    for left < right {
        mid := left + (right-left)/2  // avoid overflow
        if nums[mid] < nums[mid+1] {
            // Ascending slope: peak must be in the right half
            left = mid + 1
        } else {
            // Descending slope or at peak: peak is at mid or in the left half
            right = mid
        }
    }
    // left == right: converged to a peak
    return left
}