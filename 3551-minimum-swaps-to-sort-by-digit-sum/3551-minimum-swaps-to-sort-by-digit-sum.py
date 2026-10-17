class Solution:
    def minSwaps(self, nums: List[int]) -> int:
        # Time: O(n log n) | Space: O(n)

        def digit_sum(x: int) -> int:
            s = 0
            while x:
                s += x % 10
                x //= 10
            return s

        # Sort by (digit_sum, value)
        sorted_nums = sorted(nums, key=lambda x: (digit_sum(x), x))

        # Map value → target index
        pos = {v: i for i, v in enumerate(sorted_nums)}

        visited = [False] * len(nums)
        swaps   = 0

        for i in range(len(nums)):
            if visited[i] or nums[i] == sorted_nums[i]:
                continue

            # Trace cycle starting at i
            cycle_len = 0
            j = i
            while not visited[j]:
                visited[j] = True
                j = pos[nums[j]]  # where does nums[j] need to go?
                cycle_len += 1

            swaps += cycle_len - 1

        return swaps