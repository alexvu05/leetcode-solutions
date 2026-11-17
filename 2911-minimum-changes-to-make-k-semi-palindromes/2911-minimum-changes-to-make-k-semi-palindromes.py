class Solution:
    def minimumChanges(self, s: str, k: int) -> int:
        # Time: O(n^3 * d(n)) preprocessing (small constants) + O(k * n^2) DP
        # Space: O(n^2)
        n = len(s)
        INF = float('inf')

        # divs[length] = proper divisors d (1 <= d < length) of length
        divs = [[] for _ in range(n + 1)]
        for d in range(1, n):
            for length in range(2 * d, n + 1, d):
                divs[length].append(d)

        # cost[i][j] = min changes to make s[i:j] a semi-palindrome (j - i >= 2)
        cost = [[0] * (n + 1) for _ in range(n + 1)]
        for i in range(n):
            for j in range(i + 2, n + 1):
                best = INF
                for d in divs[j - i]:
                    mism = 0
                    for r in range(d):                  # one residue class at a time
                        l, rr = i + r, j - d + r        # first / last index of the class
                        while l < rr:
                            if s[l] != s[rr]:
                                mism += 1               # this mirrored pair must be fixed
                            l += d
                            rr -= d
                    if mism < best:
                        best = mism
                    if best == 0:
                        break                           # cannot do better
                cost[i][j] = best

        # f_prev[i] = best cost to split first i chars into p-1 pieces
        f_prev = [INF] * (n + 1)
        f_prev[0] = 0
        for p in range(1, k + 1):
            f_cur = [INF] * (n + 1)
            for i in range(2 * p, n + 1):               # p pieces need at least 2p chars
                best = INF
                for h in range(2 * (p - 1), i - 1):     # last piece s[h:i] has length >= 2
                    v = f_prev[h] + cost[h][i]
                    if v < best:
                        best = v
                f_cur[i] = best
            f_prev = f_cur

        return f_prev[n]