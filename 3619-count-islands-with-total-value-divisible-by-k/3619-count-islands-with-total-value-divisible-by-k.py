class Solution:
    def countIslands(self, grid: List[List[int]], k: int) -> int:
        # Time: O(m*n) | Space: O(m*n) worst case for the stack
        m, n = len(grid), len(grid[0])
        count = 0

        for sr in range(m):
            for sc in range(n):
                if grid[sr][sc] == 0:
                    continue

                # New island found: flood-fill and sum its cells
                total = 0
                stack = [(sr, sc)]
                total += grid[sr][sc]
                grid[sr][sc] = 0                 # mark visited when pushing

                while stack:
                    r, c = stack.pop()
                    for dr, dc in ((1, 0), (-1, 0), (0, 1), (0, -1)):
                        nr, nc = r + dr, c + dc
                        if 0 <= nr < m and 0 <= nc < n and grid[nr][nc] > 0:
                            total += grid[nr][nc]
                            grid[nr][nc] = 0     # mark visited
                            stack.append((nr, nc))

                if total % k == 0:
                    count += 1

        return count