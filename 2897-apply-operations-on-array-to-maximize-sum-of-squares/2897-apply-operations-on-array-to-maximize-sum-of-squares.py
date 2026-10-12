class Solution:
    def maxSum(self, nums: List[int], k: int) -> int:
        # Time: O(n * 30) | Space: O(30)
        MOD = 10**9 + 7

        # Count occurrences of each bit across all numbers
        cnt = [0] * 30  # bits 0..29
        for x in nums:
            for b in range(30):
                if x >> b & 1:
                    cnt[b] += 1

        # Build k optimal numbers greedily
        # Number j (0-indexed) gets bit b if j < cnt[b]
        # i.e., the first cnt[b] numbers get bit b
        # So number j = sum of (1<<b) for all b where cnt[b] > j

        result = 0
        for j in range(k):
            val = 0
            for b in range(30):
                if cnt[b] > j:
                    val |= (1 << b)
            result = (result + val * val) % MOD

        return result