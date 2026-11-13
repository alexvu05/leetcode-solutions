class Solution:
    def countNonDecreasingSubarrays(self, nums: List[int], k: int) -> int:
        # Time: O(n) amortized | Space: O(n)
        n = len(nums)
        ans = 0
        cost = 0                  # cost to make nums[left..right] non-decreasing
        dq = deque()              # prefix-max "records" of the window; back = left, front = farthest
        right = n - 1

        for left in range(n - 1, -1, -1):
            # Add nums[left]: absorb smaller records, raising their segments up to nums[left]
            while dq and nums[dq[-1]] < nums[left]:
                l = dq.pop()
                r = dq[-1] - 1 if dq else right      # last index covered by record l
                cost += (r - l + 1) * (nums[left] - nums[l])
            dq.append(left)

            # Shrink from the right while the window is too expensive
            while cost > k:
                cost -= nums[dq[0]] - nums[right]    # running max at `right` is nums[front record]
                if dq[0] == right:
                    dq.popleft()                     # that record covered only `right`
                right -= 1

            ans += right - left + 1                  # every subarray left..r with r <= right works

        return ans