class Solution:
    def closedIsland(self, grid: list[list[int]]) -> int:
        # Time: O(m*n) | Space: O(m*n) worst case for stack
        m, n = len(grid), len(grid[0])

        def flood(sr: int, sc: int) -> None:
            # Turn the whole island containing (sr, sc) into water
            stack = [(sr, sc)]
            grid[sr][sc] = 1                      # mark when pushing → no duplicates
            while stack:
                r, c = stack.pop()
                for dr, dc in ((1, 0), (-1, 0), (0, 1), (0, -1)):
                    nr, nc = r + dr, c + dc
                    if 0 <= nr < m and 0 <= nc < n and grid[nr][nc] == 0:
                        grid[nr][nc] = 1
                        stack.append((nr, nc))

        # Step 1: remove every island that touches the border
        for r in range(m):
            for c in range(n):
                on_border = r in (0, m - 1) or c in (0, n - 1)
                if on_border and grid[r][c] == 0:
                    flood(r, c)

        # Step 2: remaining islands are all closed → count them
        count = 0
        for r in range(1, m - 1):
            for c in range(1, n - 1):
                if grid[r][c] == 0:
                    count += 1
                    flood(r, c)

        return count