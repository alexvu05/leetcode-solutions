class Solution:
    def removeOuterParentheses(self, s: str) -> str:
        # Time: O(n) | Space: O(n)
        result = []
        depth  = 0

        for c in s:
            if c == '(':
                if depth > 0:       # not the outermost '('
                    result.append(c)
                depth += 1
            else:
                depth -= 1
                if depth > 0:       # not the outermost ')'
                    result.append(c)

        return "".join(result)