class Solution:
    def createGrid(self, m: int, n: int) -> list[str]:
        # Time: O(m*n) to build the output | Space: O(m*n) for the output
        first_row = '.' * n                 # row 0: all free
        other_row = '#' * (n - 1) + '.'     # rows 1..m-1: only last column free

        return [first_row] + [other_row] * (m - 1)