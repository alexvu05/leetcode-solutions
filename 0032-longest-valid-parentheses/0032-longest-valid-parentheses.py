class Solution:
    def longestValidParentheses(self, s: str) -> int:
        # Stack stores indices; bottom is always the last unmatched ')' index
        stack = [-1]  # sentinel: index -1 as base
        best = 0
        for i, c in enumerate(s):
            if c == '(':
                stack.append(i)
            else:
                stack.pop()
                if not stack:
                    # No unmatched '(' → this ')' becomes new base/barrier
                    stack.append(i)
                else:
                    # Distance from current to last unmatched = valid length
                    best = max(best, i - stack[-1])
        return best