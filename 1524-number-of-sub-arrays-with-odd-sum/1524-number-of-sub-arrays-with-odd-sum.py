class Solution:
    def numOfSubarrays(self, arr: list[int]) -> int:
        # Time: O(n) | Space: O(1)
        MOD = 10**9 + 7
        cnt = [1, 0]        # cnt[p] = number of prefixes with parity p; empty prefix is even
        p = 0               # parity of the current prefix sum
        ans = 0

        for x in arr:
            p ^= x & 1                  # flip parity iff x is odd
            ans = (ans + cnt[p ^ 1]) % MOD   # pair with earlier prefixes of opposite parity
            cnt[p] += 1

        return ans