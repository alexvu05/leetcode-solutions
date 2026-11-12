class Solution:
    def longestMonotonicSubarray(self, nums: List[int]) -> int:
        # Time: O(n) | Space: O(1)
        best = inc = dec = 1          # a single element is always a valid subarray

        for i in range(1, len(nums)):
            if nums[i] > nums[i - 1]:
                inc += 1              # extend increasing run
                dec = 1               # decreasing run is broken
            elif nums[i] < nums[i - 1]:
                dec += 1              # extend decreasing run
                inc = 1               # increasing run is broken
            else:
                inc = dec = 1         # equal neighbors break both runs

            best = max(best, inc, dec)

        return best