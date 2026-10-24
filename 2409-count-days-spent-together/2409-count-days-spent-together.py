class Solution:
    def countDaysTogether(self, arriveAlice: str, leaveAlice: str, arriveBob: str, leaveBob: str) -> int:
        # Time: O(1) | Space: O(1)
        DAYS = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]  # non-leap year

        def to_day(s: str) -> int:
            month, day = int(s[:2]), int(s[3:])
            return sum(DAYS[:month - 1]) + day   # day-of-year

        start = max(to_day(arriveAlice), to_day(arriveBob))
        end   = min(to_day(leaveAlice),  to_day(leaveBob))

        return max(0, end - start + 1)