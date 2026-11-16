class Solution:
    def digitSum(self, s: str, k: int) -> str:
        # Time: O(n) per round, few rounds | Space: O(n)
        while len(s) > k:
            groups = (s[i:i + k] for i in range(0, len(s), k))     # last group may be shorter
            s = "".join(str(sum(int(c) for c in g)) for g in groups)
        return s