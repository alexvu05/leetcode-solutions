class Solution:
    def checkString(self, s: str) -> bool:
        # Time: O(n) | Space: O(1)
        seen_b = False
        for c in s:
            if c == 'b':
                seen_b = True
            elif seen_b:        # c == 'a' after a 'b' has already appeared
                return False
        return True