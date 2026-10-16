class Solution:
    def sumCounts(self, nums: List[int]) -> int:
        # Time: O(n log n) | Space: O(n)
        MOD = 10**9 + 7
        n   = len(nums)

        # Segment tree: each node stores sum of cnt[i]^2 and sum of cnt[i]
        # for range [l, r], with lazy add
        size   = 1
        while size < n: size <<= 1

        sum_sq  = [0] * (2 * size)  # sum of cnt[i]^2
        sum_cnt = [0] * (2 * size)  # sum of cnt[i]
        lazy    = [0] * (2 * size)  # pending add

        def push_down(node, l, r):
            if lazy[node] == 0: return
            mid = (l + r) // 2
            add(2*node,   l,   mid, lazy[node])
            add(2*node+1, mid+1, r, lazy[node])
            lazy[node] = 0

        def add(node, l, r, val):
            # Add val to all cnt[i] in [l, r]
            length = r - l + 1
            # sum(cnt[i]+val)^2 = sum(cnt[i]^2) + 2*val*sum(cnt[i]) + val^2*length
            sum_sq[node]  = (sum_sq[node] + 2 * val * sum_cnt[node] + val * val * length) % MOD
            sum_cnt[node] = (sum_cnt[node] + val * length) % MOD
            lazy[node]    = (lazy[node] + val) % MOD

        def update(node, l, r, ql, qr):
            if ql > r or qr < l: return
            if ql <= l and r <= qr:
                add(node, l, r, 1)
                return
            push_down(node, l, r)
            mid = (l + r) // 2
            update(2*node,   l,   mid, ql, qr)
            update(2*node+1, mid+1, r, ql, qr)
            sum_sq[node]  = (sum_sq[2*node] + sum_sq[2*node+1]) % MOD
            sum_cnt[node] = (sum_cnt[2*node] + sum_cnt[2*node+1]) % MOD

        ans  = 0
        last = {}  # last seen index of each value

        for j in range(n):
            l = last.get(nums[j], -1) + 1  # range [l, j] gets +1 distinct
            update(1, 0, size-1, l, j)
            ans = (ans + sum_sq[1]) % MOD
            last[nums[j]] = j

        return ans