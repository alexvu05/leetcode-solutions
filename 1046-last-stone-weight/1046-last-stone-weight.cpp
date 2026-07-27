class Solution {
public:
    int lastStoneWeight(vector<int>& stones) {
        // Max-heap: largest stone always at top
        priority_queue<int> maxHeap(stones.begin(), stones.end());
        while (maxHeap.size() > 1) {
            int y = maxHeap.top(); maxHeap.pop();  // heaviest
            int x = maxHeap.top(); maxHeap.pop();  // second heaviest
            // Only push remainder if stones differ
            if (y != x) {
                maxHeap.push(y - x);
            }
        }
        return maxHeap.empty() ? 0 : maxHeap.top();
    }
};