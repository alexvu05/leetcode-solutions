class Solution:
    def minOperations(self, s: str) -> int:
        # Time: O(n) | Space: O(1)
        n = len(s)

        # Case 1: already sorted
        if all(s[i] <= s[i + 1] for i in range(n - 1)):
            return 0

        # Case 2: length 2, unsorted → no proper substring can fix it
        if n == 2:
            return -1

        mn, mx = min(s), max(s)

        # Case 3: one end is already correct → sort the rest
        if s[0] == mn or s[-1] == mx:
            return 1

        # Case 4: min is somewhere not-last, or max is somewhere not-first
        if s.find(mn) != n - 1 or s.rfind(mx) != 0:
            return 2

        # Case 5: min only at the last position AND max only at the first
        return 3