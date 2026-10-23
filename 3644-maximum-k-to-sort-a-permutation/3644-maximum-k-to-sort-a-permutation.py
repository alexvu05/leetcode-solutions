class Solution:
    def sortPermutation(self, nums: List[int]) -> int:
        # Time: O(n) | Space: O(1)
        ans = -1   # all bits set: identity element for AND

        for i, x in enumerate(nums):
            if x != i:          # misplaced element
                ans &= x

        # If nothing was misplaced, ans is still -1 → return 0
        return max(ans, 0)