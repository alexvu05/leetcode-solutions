class Solution:
    def minMoves(self, nums: list[int]) -> int:
        # Time: O(n) | Space: O(1)
        return sum(nums) - len(nums) * min(nums)