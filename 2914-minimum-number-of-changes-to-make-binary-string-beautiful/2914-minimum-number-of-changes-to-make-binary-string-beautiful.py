class Solution:
    def minChanges(self, s: str) -> int:
        # Time: O(n) | Space: O(1)
        # Count aligned pairs (s[0],s[1]), (s[2],s[3]), ... whose two characters differ
        return sum(s[i] != s[i + 1] for i in range(0, len(s), 2))