class Solution:
    def minMoves(self, nums: List[int]) -> int:
        # Time: O(n) | Space: O(1)
        return len(nums) * max(nums) - sum(nums)