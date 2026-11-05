class Solution:
    def findTheDistanceValue(self, arr1: list[int], arr2: list[int], d: int) -> int:
        # Time: O((n + m) log m) | Space: O(1) extra (sort in place)
        arr2.sort()
        m = len(arr2)
        count = 0

        for x in arr1:
            i = bisect_left(arr2, x - d)          # first element >= x - d
            # Valid if no such element, or it already lies beyond x + d
            if i == m or arr2[i] > x + d:
                count += 1

        return count