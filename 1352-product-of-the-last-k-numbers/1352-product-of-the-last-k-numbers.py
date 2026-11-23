class ProductOfNumbers:
    # add: O(1) | getProduct: O(1) | Space: O(n)

    def __init__(self):
        self.prefix = [1]                  # prefix[i] = product of the first i numbers after the last 0

    def add(self, num: int) -> None:
        if num == 0:
            self.prefix = [1]              # any window touching this 0 has product 0 → start over
        else:
            self.prefix.append(self.prefix[-1] * num)

    def getProduct(self, k: int) -> int:
        if k >= len(self.prefix):          # only len(prefix)-1 numbers since the last 0
            return 0                       # window of size k reaches the zero
        return self.prefix[-1] // self.prefix[-1 - k]
        


# Your ProductOfNumbers object will be instantiated and called as such:
# obj = ProductOfNumbers()
# obj.add(num)
# param_2 = obj.getProduct(k)