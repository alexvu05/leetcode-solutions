class Solution:
    def countServers(self, grid: list[list[int]]) -> int:
        # Time: O(m*n) | Space: O(m + n)
        m, n = len(grid), len(grid[0])
        row = [0] * m          # number of servers in each row
        col = [0] * n          # number of servers in each column

        for i in range(m):
            for j in range(n):
                if grid[i][j]:
                    row[i] += 1
                    col[j] += 1

        # A server communicates if its row OR its column holds another server
        return sum(
            1
            for i in range(m)
            for j in range(n)
            if grid[i][j] and (row[i] > 1 or col[j] > 1)
        )