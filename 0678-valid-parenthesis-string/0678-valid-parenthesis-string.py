class Solution:
    def checkValidString(self, s: str) -> bool:
        lo = hi = 0  # [lo, hi] = range of possible balances
        for c in s:
            if c == '(':
                lo += 1
                hi += 1
            elif c == ')':
                lo -= 1
                hi -= 1
            else:  # '*'
                lo -= 1  # treat '*' as ')'
                hi += 1  # treat '*' as '('
            # Balance can't go negative
            if hi < 0:
                return False  # too many ')' even with all '*' as '('
            lo = max(lo, 0)   # clamp: can't have negative balance
        # Valid if balance 0 is achievable
        return lo == 0