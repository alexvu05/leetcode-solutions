class Solution {
public:
    int maxNumberOfFamilies(int n, vector<vector<int>>& reservedSeats) {
        // Bitmask: seat s maps to bit (s-1), seats 1-10 → bits 0-9
        // Valid blocks (seats are 1-indexed):
        // LEFT  = seats 2,3,4,5 → bits 1,2,3,4 → mask = 0b0000011110 = 30
        // MID   = seats 4,5,6,7 → bits 3,4,5,6 → mask = 0b0001111000 = ... 
        // Let's use seat s → bit (s-1):
        // LEFT  seats 2,3,4,5 → bits 1,2,3,4 → (1<<1)|(1<<2)|(1<<3)|(1<<4) = 30
        // MID   seats 4,5,6,7 → bits 3,4,5,6 → (1<<3)|(1<<4)|(1<<5)|(1<<6) = 120
        // RIGHT seats 6,7,8,9 → bits 5,6,7,8 → (1<<5)|(1<<6)|(1<<7)|(1<<8) = 480
        const int LEFT  = (1<<1)|(1<<2)|(1<<3)|(1<<4);  // seats 2,3,4,5
        const int MID   = (1<<3)|(1<<4)|(1<<5)|(1<<6);  // seats 4,5,6,7
        const int RIGHT = (1<<5)|(1<<6)|(1<<7)|(1<<8);  // seats 6,7,8,9
        // Build reserved bitmask per row (ignore seats 1 and 10 — not in any block)
        unordered_map<int, int> rowMask;
        for (auto& seat : reservedSeats) {
            int row = seat[0], col = seat[1];
            if (col == 1 || col == 10) continue;  // irrelevant seats
            rowMask[row] |= (1 << (col - 1));
        }
        // Rows with no reservations: always fit 2 groups
        int result = 2 * (n - (int)rowMask.size());
        // Process rows with reservations
        for (auto& [row, mask] : rowMask) {
            bool canLeft  = (mask & LEFT)  == 0;
            bool canMid   = (mask & MID)   == 0;
            bool canRight = (mask & RIGHT) == 0;
            if (canLeft && canRight) {
                result += 2;  // both sides free
            } else if (canLeft || canMid || canRight) {
                result += 1;  // only one block available
            }
            // else: 0 groups for this row
        }
        return result;
    }
};