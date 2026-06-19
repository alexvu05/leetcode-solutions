class Solution {
public:
    int singleNumber(vector<int>& nums) {
        int result = 0;
        // XOR-ing a number with itself cancels it out (a ^ a = 0);
        // the number appearing once survives since x ^ 0 = x
        for (int x : nums) {
            result ^= x;
        }
        return result;
    }
};