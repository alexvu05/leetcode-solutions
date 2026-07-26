class Solution {
public:
    vector<vector<int>> kClosest(vector<vector<int>>& points, int k) {
        // Max-heap: {squared_distance, index}
        // Largest distance stays at top → easy to evict when heap exceeds k
        priority_queue<pair<int,int>> maxHeap;
        for (int i = 0; i < (int)points.size(); i++) {
            int dist = points[i][0]*points[i][0] + points[i][1]*points[i][1];
            maxHeap.push({dist, i});
            // Evict the farthest point when heap exceeds k
            if ((int)maxHeap.size() > k) {
                maxHeap.pop();
            }
        }
        // Collect the k closest points
        vector<vector<int>> result;
        result.reserve(k);
        while (!maxHeap.empty()) {
            result.push_back(points[maxHeap.top().second]);
            maxHeap.pop();
        }
        return result;
    }
};