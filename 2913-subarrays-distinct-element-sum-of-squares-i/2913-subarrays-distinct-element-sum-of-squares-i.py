class Solution:
    def sumCounts(self, nums: List[int]) -> int:
        # Time: O(n^2) | Space: O(n)
        n      = len(nums)
        result = 0

        for i in range(n):
            seen = set()
            for j in range(i, n):
                seen.add(nums[j])
                result += len(seen) ** 2

        return result