class Solution:
    def sumOfSquares(self, nums: List[int]) -> int:
        # Time: O(n) | Space: O(1)
        n      = len(nums)
        result = 0

        for i in range(1, n + 1):    # 1-indexed
            if n % i == 0:           # i divides n → special
                result += nums[i-1] * nums[i-1]

        return result