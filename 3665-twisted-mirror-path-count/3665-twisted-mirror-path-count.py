class Solution:
    def uniquePaths(self, grid: List[List[int]]) -> int:
        # Time: O(m*n) | Space: O(n)
        MOD = 10**9 + 7
        m, n = len(grid), len(grid[0])

        # out_down[c]: ways to leave the cell above (current column) heading down.
        # Virtual cell above (0,0) gives exactly one way to "enter" the start.
        out_down = [0] * n
        out_down[0] = 1

        out_right = 0   # ways to leave the previous cell in this row heading right

        for r in range(m):
            out_right = 0                    # nothing comes from the left of column 0
            for c in range(n):
                from_above = out_down[c]     # entered moving down
                from_left = out_right        # entered moving right

                if grid[r][c] == 1:          # mirror: swap directions
                    new_right, new_down = from_above, from_left
                else:                        # empty: can leave either way
                    new_right = new_down = (from_above + from_left) % MOD

                out_down[c] = new_down
                out_right = new_right

        return out_right                     # ways to leave last cell heading right