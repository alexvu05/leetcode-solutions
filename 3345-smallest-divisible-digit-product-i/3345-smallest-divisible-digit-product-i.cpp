class Solution {
public:
    int smallestNumber(int n, int t) {
        for (int num = n; ; num++) {
            if (digitProduct(num) % t == 0) {
                return num;
            }
        }
    }
private:
    // Compute product of all digits of x
    int digitProduct(int x) {
        int product = 1;
        while (x > 0) {
            product *= x % 10; // extract last digit
            x /= 10; // remove last digit
        }
        return product;
    }
};
