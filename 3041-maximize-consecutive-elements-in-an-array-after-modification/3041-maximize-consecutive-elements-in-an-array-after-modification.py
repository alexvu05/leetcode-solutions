class Solution:
    def maxSelectedElements(self, nums: List[int]) -> int:
        # Time: O(n log n) | Space: O(n)
        dp = defaultdict(int)        # dp[v] = longest consecutive run ending at value v

        for x in sorted(nums):
            dp[x + 1] = dp[x] + 1    # raise this element to x+1  (read dp[x] BEFORE updating it)
            dp[x] = dp[x - 1] + 1    # keep this element at x

        return max(dp.values())