class Solution:
    def generateParenthesis(self, n: int) -> list[str]:
        result = []
        def backtrack(current: str, open: int, close: int) -> None:
            # Base case: used all n pairs
            if open == n and close == n:
                result.append(current)
                return
            # Add '(' if we haven't used all n open brackets
            if open < n:
                backtrack(current + "(", open + 1, close)
            # Add ')' only if there's an unmatched '(' to close
            if close < open:
                backtrack(current + ")", open, close + 1)
        backtrack("", 0, 0)
        return result