class Solution:
    def climbStairs(self, n: int, costs: List[int]) -> int:
        # Time: O(n) | Space: O(n)  (O(1) with 3 rolling variables)
        INF = float('inf')
        f = [INF] * (n + 1)
        f[0] = 0

        for i in range(1, n + 1):
            best = INF
            for jump in (1, 2, 3):
                j = i - jump
                if j >= 0:                              # step must exist
                    best = min(best, f[j] + jump * jump)
            f[i] = best + costs[i - 1]                  # costs is 1-indexed in the statement

        return f[n]