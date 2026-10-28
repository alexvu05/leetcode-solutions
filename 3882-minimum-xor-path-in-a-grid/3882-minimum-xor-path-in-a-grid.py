class Solution:
    def minCost(self, grid: list[list[int]]) -> int:
        import numpy as np
        # Time: O(m*n*16*6) element-ops, vectorized over anti-diagonals
        # Space: O(min(m,n) * 16)
        m, n = len(grid), len(grid[0])
        g = np.array(grid, dtype=np.int64)

        # MASK[b]: positions in a 64-bit word whose bit b is 0
        MASK = [np.uint64(x) for x in (
            0x5555555555555555, 0x3333333333333333, 0x0F0F0F0F0F0F0F0F,
            0x00FF00FF00FF00FF, 0x0000FFFF0000FFFF, 0x00000000FFFFFFFF)]
        W = np.arange(16, dtype=np.int64)

        prev = None
        for d in range(m + n - 1):
            lo, hi = max(0, d - n + 1), min(d, m - 1)       # rows on this diagonal
            cnt = hi - lo + 1

            merged = np.zeros((cnt, 16), dtype=np.uint64)
            if d == 0:
                merged[0, 0] = 1                             # XOR so far = 0
            else:
                lo_p, hi_p = max(0, d - n), min(d - 1, m - 1)  # rows on diagonal d-1
                # from the left: same row i
                a, b = max(lo, lo_p), min(hi, hi_p)
                if a <= b:
                    merged[a - lo:b - lo + 1] |= prev[a - lo_p:b - lo_p + 1]
                # from above: row i-1
                a, b = max(lo, lo_p + 1), min(hi, hi_p + 1)
                if a <= b:
                    merged[a - lo:b - lo + 1] |= prev[a - 1 - lo_p:b - lo_p]

            # Apply v -> v ^ g for every cell of the diagonal
            ii = np.arange(lo, hi + 1)
            gv = g[ii, d - ii]
            gh, gl = gv >> 6, gv & 63

            # permute the 16 words: res[w] = merged[w ^ gh]
            idx = W[None, :] ^ gh[:, None]
            x = np.take_along_axis(merged, idx, axis=1)

            # permute bits inside each word, only where bit b of gl is set
            for b in range(6):
                sel = ((gl >> b) & 1).astype(bool)
                if sel.any():
                    s, mb = np.uint64(1 << b), MASK[b]
                    swapped = ((x & mb) << s) | ((x >> s) & mb)
                    x = np.where(sel[:, None], swapped, x)

            prev = x

        # last diagonal has exactly one cell: (m-1, n-1)
        for w in range(16):
            word = int(prev[0, w])
            if word:
                return w * 64 + (word & -word).bit_length() - 1
        return -1