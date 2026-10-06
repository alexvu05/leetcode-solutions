class Solution:
    def minAddToMakeValid(self, s: str) -> int:
        open  = 0  # unmatched '('
        close = 0  # unmatched ')'
        for c in s:
            if c == '(':
                open += 1
            else:
                if open > 0:
                    open -= 1  # match with existing '('
                else:
                    close += 1  # no '(' to match → need to add one
        # open = unmatched '(' needing ')' added
        # close = unmatched ')' needing '(' added
        return open + close