class Solution:
    def targetIndices(self, nums: list[int], target: int) -> list[int]:
        # Time: O(n) | Space: O(1) excluding output
        less = equal = 0

        for x in nums:
            if x < target:
                less += 1
            elif x == target:
                equal += 1

        # Sorted positions of target: a contiguous block starting at 'less'
        return list(range(less, less + equal))