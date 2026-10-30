class Solution:
    def shortestPath(self, n: int, edges: List[List[int]], labels: str, k: int) -> int:
        # Time: O((E * k) log(E * k)) worst case, much less in practice
        # Space: O(E * k) worst case for the heap
        adj = [[] for _ in range(n)]
        for u, v, w in edges:
            adj[u].append((v, w))

        # best_run[u] = smallest run length already settled at node u
        best_run = [k + 1] * n
        heap = [(0, 0, 1)]                       # (dist, node, run)

        while heap:
            d, u, r = heapq.heappop(heap)

            # Dominated: an earlier pop had dist <= d AND run <= r
            if r >= best_run[u]:
                continue
            best_run[u] = r

            if u == n - 1:
                return d                         # first non-dominated pop = shortest

            for v, w in adj[u]:
                nr = r + 1 if labels[v] == labels[u] else 1
                if nr <= k and nr < best_run[v]:
                    heapq.heappush(heap, (d + w, v, nr))

        return -1