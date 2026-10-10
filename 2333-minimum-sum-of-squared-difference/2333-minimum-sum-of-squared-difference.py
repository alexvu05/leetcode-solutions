class Solution:
    def minSumSquareDiff(self, nums1: list[int], nums2: list[int], k1: int, k2: int) -> int:
        # Time: O(n log n + n log(maxVal)) | Space: O(n)
        k    = k1 + k2
        diff = sorted(abs(a - b) for a, b in zip(nums1, nums2))
        n    = len(diff)

        # Binary search: find smallest T such that
        # total ops to reduce all diff[i] to min(diff[i], T) <= k
        lo, hi = 0, diff[-1]
        while lo < hi:
            mid  = (lo + hi) // 2
            cost = sum(max(0, d - mid) for d in diff)
            if cost <= k:
                hi = mid
            else:
                lo = mid + 1

        T = lo

        # Ops used to bring everything to T
        used      = sum(max(0, d - T) for d in diff)
        remaining = k - used  # leftover ops reduce some T→(T-1)

        result = 0
        for d in diff:
            val = min(d, T)
            if val == T and remaining > 0 and T > 0:
                val      -= 1
                remaining -= 1
            result += val * val

        return result