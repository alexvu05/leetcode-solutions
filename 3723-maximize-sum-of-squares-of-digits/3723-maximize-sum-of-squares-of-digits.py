class Solution:
    def maxSumOfSquares(self, num: int, sum: int) -> str:
        # Time: O(num) | Space: O(num)

        # Impossible: can't form num digits with digit sum = sum
        if sum > 9 * num:
            return ""

        k = sum // 9   # number of 9s
        r = sum % 9    # remainder digit (0 means no remainder digit needed)

        # Build: k nines, then remainder (if any), then zeros
        # This gives the lexicographically/numerically largest number with max score
        result = ["9"] * k

        if r > 0:
            result.append(str(r))

        # Pad with zeros to reach exactly num digits
        result += ["0"] * (num - len(result))

        return "".join(result)