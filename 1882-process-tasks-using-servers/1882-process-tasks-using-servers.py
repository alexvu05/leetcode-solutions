class Solution:
    def assignTasks(self, servers: list[int], tasks: list[int]) -> list[int]:
        # Time: O((n + m) log n) | Space: O(n + m)
        free = [(w, i) for i, w in enumerate(servers)]   # (weight, index)
        heapq.heapify(free)
        busy = []                                        # (finish_time, weight, index)

        ans = [0] * len(tasks)
        t = 0

        for j, dur in enumerate(tasks):
            t = max(t, j)                    # task j exists only from second j

            # Release every server that has finished by time t
            while busy and busy[0][0] <= t:
                _, w, i = heapq.heappop(busy)
                heapq.heappush(free, (w, i))

            # Nobody free: jump to the earliest finish time
            if not free:
                t = busy[0][0]
                while busy and busy[0][0] <= t:
                    _, w, i = heapq.heappop(busy)
                    heapq.heappush(free, (w, i))

            w, i = heapq.heappop(free)       # lowest weight, then lowest index
            ans[j] = i
            heapq.heappush(busy, (t + dur, w, i))

        return ans