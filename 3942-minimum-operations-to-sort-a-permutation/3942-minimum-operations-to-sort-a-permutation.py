class Solution:
    def minOperations(self, nums: List[int]) -> int:
        # Time: O(n) | Space: O(1)
        n = len(nums)
        if n == 1:
            return 0

        z = nums.index(0)   # position of 0

        # Is nums a rotation of [0,1,...,n-1] ?
        inc = all(nums[(i + 1) % n] == (nums[i] + 1) % n for i in range(n))
        # Is nums a rotation of [n-1,...,1,0] ?
        dec = all(nums[(i + 1) % n] == (nums[i] - 1) % n for i in range(n))

        INF = float('inf')
        ans = INF

        if inc:
            # rotate left z times, or (reverse, rotate, reverse) = rotate right n-z
            ans = min(ans, z, n - z + 2)
        if dec:
            # reverse then rotate left (n-1-z), or rotate left z+1 then reverse
            ans = min(ans, n - z, z + 2)

        return -1 if ans == INF else ans