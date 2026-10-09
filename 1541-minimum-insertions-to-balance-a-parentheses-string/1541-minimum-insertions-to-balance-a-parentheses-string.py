class Solution:
    def minInsertions(self, s: str) -> int:
        # Time: O(n) | Space: O(1)
        res  = 0  # insertions needed
        open = 0  # unmatched '(' (each needs 2 ')')
        i    = 0

        while i < len(s):
            c = s[i]
            if c == '(':
                open += 1
                i += 1
            else:  # c == ')'
                # Check if next char is also ')'
                if i + 1 < len(s) and s[i+1] == ')':
                    # Got "))" → consume one open '('
                    i += 2
                else:
                    # Only one ')' → need to insert one more ')'
                    res += 1
                    i += 1

                # Now we have one "))" (real or with insertion)
                if open > 0:
                    open -= 1  # match with existing '('
                else:
                    # No '(' to match → need to insert '('
                    res += 1

        # Each remaining open '(' needs "))"
        res += open * 2

        return res