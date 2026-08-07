class Solution {
public:
    int minimumCost(vector<int>& cost) {
        // Sort descending: most expensive first
        sort(cost.begin(), cost.end(), greater<int>());
        int total = 0;
        for (int i = 0; i < (int)cost.size(); i++) {
            // Every 3rd item (0-indexed: 2,5,8,...) is free
            if ((i + 1) % 3 == 0) continue;
            total += cost[i];
        }
        return total;
    }
};