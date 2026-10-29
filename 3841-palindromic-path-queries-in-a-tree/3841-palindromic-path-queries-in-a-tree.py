class Solution:
    def palindromePath(self, n: int, edges: list[list[int]], s: str, queries: list[str]) -> list[bool]:
        # Time: O((n + q) log n) | Space: O(n log n)
        adj = [[] for _ in range(n)]
        for a, b in edges:
            adj[a].append(b)
            adj[b].append(a)

        bit = [1 << (ord(ch) - 97) for ch in s]   # current char of each node as a bit

        # ---- iterative DFS: preorder (tin), parent, depth, base prefix-XOR ----
        parent = [0] * n
        depth = [0] * n
        base = [0] * n            # P[x] with the ORIGINAL characters
        tin = [0] * n
        order = []
        base[0] = bit[0]
        stack = [0]
        while stack:
            u = stack.pop()
            tin[u] = len(order)
            order.append(u)
            for v in adj[u]:
                if v != parent[u] or (u == 0 and v != 0):
                    if v == parent[u] and u != 0:
                        continue
                    parent[v] = u
                    depth[v] = depth[u] + 1
                    base[v] = base[u] ^ bit[v]
                    stack.append(v)

        # subtree sizes -> tout (subtree of u = [tin[u], tout[u]])
        size = [1] * n
        for u in reversed(order):
            if u != 0:
                size[parent[u]] += size[u]
        tout = [tin[u] + size[u] - 1 for u in range(n)]

        # ---- binary lifting for LCA (root's parent = itself) ----
        LOG = max(1, n.bit_length())
        up = [parent[:]]
        for k in range(1, LOG):
            prev = up[-1]
            up.append([prev[prev[v]] for v in range(n)])

        def lca(u: int, v: int) -> int:
            if depth[u] < depth[v]:
                u, v = v, u
            diff = depth[u] - depth[v]
            for k in range(LOG):
                if (diff >> k) & 1:
                    u = up[k][u]
            if u == v:
                return u
            for k in range(LOG - 1, -1, -1):
                if up[k][u] != up[k][v]:
                    u, v = up[k][u], up[k][v]
            return up[0][u]

        # ---- Fenwick tree: range XOR update, point query (over tin) ----
        tree = [0] * (n + 2)

        def fen_update(i: int, x: int) -> None:
            i += 1
            while i <= n + 1:
                tree[i] ^= x
                i += i & -i

        def fen_query(i: int) -> int:
            i += 1
            r = 0
            while i > 0:
                r ^= tree[i]
                i -= i & -i
            return r

        def path_prefix(x: int) -> int:        # current P[x]
            return base[x] ^ fen_query(tin[x])

        res = []
        for q in queries:
            parts = q.split()
            if parts[0] == "update":
                u = int(parts[1])
                nb = 1 << (ord(parts[2]) - 97)
                delta = bit[u] ^ nb
                if delta:
                    fen_update(tin[u], delta)           # XOR whole subtree of u
                    fen_update(tout[u] + 1, delta)
                    bit[u] = nb
            else:
                u, v = int(parts[1]), int(parts[2])
                l = lca(u, v)
                mask = path_prefix(u) ^ path_prefix(v) ^ bit[l]
                res.append(mask & (mask - 1) == 0)      # at most one odd count
        return res