class Solution:
    def maxAlternatingSum(self, nums: List[int]) -> int:
        # Time: O(n log n) | Space: O(1)

        # Sort by squared value ascending
        nums.sort(key=lambda x: x * x)
        n = len(nums)

        # First half (smaller squares): subtracted (odd positions)
        # Second half (larger squares): added (even positions)
        s1 = sum(x * x for x in nums[:n // 2])   # to subtract
        s2 = sum(x * x for x in nums[n // 2:])   # to add

        return s2 - s1