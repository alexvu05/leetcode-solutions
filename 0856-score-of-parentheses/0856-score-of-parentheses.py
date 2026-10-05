class Solution:
    def scoreOfParentheses(self, s: str) -> int:
        stack = [0]  # current level score
        for c in s:
            if c == '(':
                stack.append(0)   # new level
            else:
                v = stack.pop()
                # '()' scores 1, '(A)' scores 2*A
                stack[-1] += max(1, 2 * v)
        return stack[0]