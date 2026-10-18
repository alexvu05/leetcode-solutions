class Solution:
    def numOfWays(self, nums: list[int]) -> int:
        # Time: O(n^2) | Space: O(n^2)
        MOD = 10**9 + 7

        def count(arr: list[int]) -> int:
            if len(arr) <= 1:
                return 1

            root  = arr[0]
            left  = [x for x in arr if x < root]
            right = [x for x in arr if x > root]

            # Ways to interleave left and right, times recursive ways
            ways = comb(len(left) + len(right), len(left))
            ways = ways * count(left) % MOD
            ways = ways * count(right) % MOD

            return ways % MOD

        # Subtract 1: exclude the original ordering
        return (count(nums) - 1) % MOD