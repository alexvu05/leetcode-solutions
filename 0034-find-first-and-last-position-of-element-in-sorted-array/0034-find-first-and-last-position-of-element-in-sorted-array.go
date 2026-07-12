func searchRange(nums []int, target int) []int {
    return []int{
        findBound(nums, target, true),   // first position
        findBound(nums, target, false),  // last position
    }
}

func findBound(nums []int, target int, findFirst bool) int {
    left, right := 0, len(nums)-1
    bound := -1
    for left <= right {
        mid := left + (right-left)/2  // avoid overflow
        if nums[mid] == target {
            bound = mid  // record this position, keep searching for tighter bound
            if findFirst {
                right = mid - 1  // search left half for an earlier occurrence
            } else {
                left = mid + 1   // search right half for a later occurrence
            }
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return bound
}