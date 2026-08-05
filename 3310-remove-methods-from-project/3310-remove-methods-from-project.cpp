class Solution {
public:
    vector<int> remainingMethods(int n, int k, vector<vector<int>>& invocations) {
        // Build adjacency list (directed: caller → callee)
        vector<vector<int>> graph(n);
        for (auto& inv : invocations) {
            graph[inv[0]].push_back(inv[1]);
        }
        // Step 1: BFS from k to find all suspicious methods
        vector<bool> suspicious(n, false);
        queue<int> bfsQueue;
        bfsQueue.push(k);
        suspicious[k] = true;
        while (!bfsQueue.empty()) {
            int node = bfsQueue.front();
            bfsQueue.pop();
            for (int callee : graph[node]) {
                if (!suspicious[callee]) {
                    suspicious[callee] = true;
                    bfsQueue.push(callee);
                }
            }
        }
        // Step 2: Check if any non-suspicious method calls a suspicious one
        for (auto& inv : invocations) {
            int caller = inv[0];
            int callee = inv[1];
            // Cross-edge: non-suspicious → suspicious → cannot remove anything
            if (!suspicious[caller] && suspicious[callee]) {
                vector<int> all(n);
                iota(all.begin(), all.end(), 0);
                return all;
            }
        }
        // Step 3: Return all non-suspicious methods
        vector<int> result;
        for (int i = 0; i < n; i++) {
            if (!suspicious[i]) {
                result.push_back(i);
            }
        }
        return result;
    }
};