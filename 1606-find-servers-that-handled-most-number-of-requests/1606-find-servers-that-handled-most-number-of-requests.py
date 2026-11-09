class Solution:
    def busiestServers(self, k: int, arrival: list[int], load: list[int]) -> list[int]:
        # Time: O(n log k) | Space: O(k)
        available = list(range(k))            # "shifted" indices of free servers
        heapq.heapify(available)
        busy = []                             # (finish_time, server)
        count = [0] * k

        for i, (a, d) in enumerate(zip(arrival, load)):
            # Release servers that have finished by the time request i arrives
            while busy and busy[0][0] <= a:
                _, s = heapq.heappop(busy)
                # smallest value >= i that is congruent to s (mod k)
                heapq.heappush(available, i + (s - i) % k)

            if not available:                 # all servers are busy → drop
                continue

            s = heapq.heappop(available) % k  # first free server at or after i % k, cyclically
            count[s] += 1
            heapq.heappush(busy, (a + d, s))

        best = max(count)
        return [s for s in range(k) if count[s] == best]