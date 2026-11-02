class Solution:
    def minMoves2(self, nums: list[int]) -> int:
        # Time: O(n log n) | Space: O(1) extra (sort in place)
        nums.sort()
        median = nums[len(nums) // 2]
        return sum(abs(x - median) for x in nums)