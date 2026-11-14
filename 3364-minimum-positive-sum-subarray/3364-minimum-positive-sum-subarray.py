class Solution:
    def minimumSumSubarray(self, nums: List[int], l: int, r: int) -> int:
        # Time: O(n * (r - l + 1)) | Space: O(n)
        n = len(nums)
        pre = [0] * (n + 1)
        for i, x in enumerate(nums):
            pre[i + 1] = pre[i] + x               # pre[i] = sum of nums[0..i-1]

        INF = float('inf')
        best = INF

        for i in range(n):
            # subarray nums[i : i+length], length in [l, r], must stay inside the array
            for length in range(l, min(r, n - i) + 1):
                s = pre[i + length] - pre[i]
                if 0 < s < best:                  # strictly positive and smaller than best
                    best = s

        return -1 if best == INF else best