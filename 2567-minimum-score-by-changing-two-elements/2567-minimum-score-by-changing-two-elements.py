class Solution:
    def minimizeSum(self, nums: List[int]) -> int:
        # Time: O(n log n) | Space: O(1) extra (sort in place)
        nums.sort()
        n = len(nums)

        # low score can always be made 0, so minimize the range [min, max]
        return min(
            nums[n - 1] - nums[2],    # change the two smallest
            nums[n - 3] - nums[0],    # change the two largest
            nums[n - 2] - nums[1],    # change one smallest and one largest
        )