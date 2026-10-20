class Solution:
    def sortableIntegers(self, nums: list[int]) -> int:
        # Time: O(n * d(n)) — d(n) là số ước của n | Space: O(1)
        n = len(nums)

        # Enumerate all divisors of n
        divisors = []
        d = 1
        while d * d <= n:
            if n % d == 0:
                divisors.append(d)
                if d != n // d:
                    divisors.append(n // d)
            d += 1

        def is_sortable(k: int) -> bool:
            prev_max = float('-inf')

            for start in range(0, n, k):
                # Count cyclic descents inside the block
                descents = 0
                for i in range(k):
                    if nums[start + i] > nums[start + (i + 1) % k]:
                        descents += 1
                        if descents > 1:
                            return False   # not a rotation of a sorted sequence

                block = nums[start:start + k]

                # After rotating, this block is sorted: [min ... max]
                # It must continue from the previous block's max
                if prev_max > min(block):
                    return False
                prev_max = max(block)

            return True

        return sum(k for k in divisors if is_sortable(k))