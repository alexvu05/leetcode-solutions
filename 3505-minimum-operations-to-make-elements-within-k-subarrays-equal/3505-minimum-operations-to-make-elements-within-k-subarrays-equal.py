class Solution:
    def minOperations(self, nums: List[int], x: int, k: int) -> int:
        # Time: O(n log n + n*k) | Space: O(n)
        n = len(nums)

        # ---- coordinate compression + Fenwick tree (count, sum) over ranks ----
        vals = sorted(set(nums))
        rank = {v: i + 1 for i, v in enumerate(vals)}
        m = len(vals)
        cnt = [0] * (m + 1)
        sm = [0] * (m + 1)
        top = 1 << (m.bit_length() - 1)          # highest power of two <= m

        def update(r: int, dc: int, dv: int) -> None:
            while r <= m:
                cnt[r] += dc
                sm[r] += dv
                r += r & -r

        kth = (x + 1) // 2                        # rank of the (lower) median in the window

        # ---- cost[s] for every window of size x, via sliding window ----
        cost = [0] * (n - x + 1)
        S = 0                                     # sum of the current window
        for i in range(n):
            update(rank[nums[i]], 1, nums[i])
            S += nums[i]
            if i >= x:                            # drop the element that left the window
                update(rank[nums[i - x]], -1, -nums[i - x])
                S -= nums[i - x]
            if i >= x - 1:
                # Fenwick descent: largest pos with prefix count < kth
                pos, rem, cl, sl, step = 0, kth, 0, 0, top
                while step:
                    nxt = pos + step
                    if nxt <= m and cnt[nxt] < rem:
                        pos = nxt
                        rem -= cnt[nxt]
                        cl += cnt[nxt]
                        sl += sm[nxt]
                    step >>= 1
                med = vals[pos]                   # median value of the window
                cost[i - x + 1] = (med * cl - sl) + (S - sl) - med * (x - cl)

        # ---- DP: choose k non-overlapping windows with minimum total cost ----
        INF = 1 << 60
        prev = [0] * (n + 1)                      # j = 0: zero windows cost nothing
        for _ in range(k):
            cur = [INF] * (n + 1)
            for i in range(x, n + 1):
                take = prev[i - x] + cost[i - x]
                cur[i] = cur[i - 1] if cur[i - 1] < take else take
            prev = cur

        return prev[n]