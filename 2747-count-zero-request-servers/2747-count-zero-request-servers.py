class Solution:
    def countServers(self, n: int, logs: List[List[int]], x: int, queries: List[int]) -> List[int]:
        # Time: O(L log L + Q log Q) | Space: O(n + Q)
        logs.sort(key=lambda p: p[1])                        # by time
        order = sorted(range(len(queries)), key=lambda i: queries[i])

        cnt = [0] * (n + 1)      # requests per server inside the window
        active = 0               # servers with at least one request in the window
        l = r = 0
        res = [0] * len(queries)

        for qi in order:
            q = queries[qi]

            # Expand right: include logs with time <= q
            while r < len(logs) and logs[r][1] <= q:
                s = logs[r][0]
                if cnt[s] == 0:
                    active += 1
                cnt[s] += 1
                r += 1

            # Shrink left: drop logs with time < q - x
            while l < r and logs[l][1] < q - x:
                s = logs[l][0]
                cnt[s] -= 1
                if cnt[s] == 0:
                    active -= 1
                l += 1

            res[qi] = n - active     # servers with no request in [q - x, q]

        return res