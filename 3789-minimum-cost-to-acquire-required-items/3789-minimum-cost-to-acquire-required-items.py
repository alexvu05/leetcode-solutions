class Solution:
    def minimumCost(self, cost1: int, cost2: int, costBoth: int, need1: int, need2: int) -> int:
        # Time: O(1) | Space: O(1)
        mn, mx = min(need1, need2), max(need1, need2)

        a = need1 * cost1 + need2 * cost2                         # only type 1 and 2
        b = mx * costBoth                                         # only type 3
        c = mn * costBoth + (need1 - mn) * cost1 + (need2 - mn) * cost2   # mix

        return min(a, b, c)