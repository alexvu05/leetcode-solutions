class Solution:
    def removeInvalidParentheses(self, s: str) -> list[str]:
        # Time: O(2^n) worst | Space: O(n)

        # Count minimum removals needed
        left_remove = right_remove = 0
        for c in s:
            if c == '(':
                left_remove += 1
            elif c == ')':
                if left_remove > 0:
                    left_remove -= 1   # matched with a '('
                else:
                    right_remove += 1  # unmatched ')'

        result = set()

        def backtrack(idx: int, path: list, open: int,
                      left_rem: int, right_rem: int) -> None:
            # Base case: processed all characters
            if idx == len(s):
                if left_rem == 0 and right_rem == 0 and open == 0:
                    result.add("".join(path))
                return

            c = s[idx]

            # Option 1: Remove current char (only if it's a bracket and we still need to remove)
            if c == '(' and left_rem > 0:
                backtrack(idx + 1, path, open, left_rem - 1, right_rem)
            elif c == ')' and right_rem > 0:
                backtrack(idx + 1, path, open, left_rem, right_rem - 1)

            # Option 2: Keep current char
            path.append(c)
            if c == '(':
                backtrack(idx + 1, path, open + 1, left_rem, right_rem)
            elif c == ')':
                # Only keep ')' if there's an unmatched '(' to close
                if open > 0:
                    backtrack(idx + 1, path, open - 1, left_rem, right_rem)
            else:
                # Letter: always keep
                backtrack(idx + 1, path, open, left_rem, right_rem)
            path.pop()

        backtrack(0, [], 0, left_remove, right_remove)
        return list(result)