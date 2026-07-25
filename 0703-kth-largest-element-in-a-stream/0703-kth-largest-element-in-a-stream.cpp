class KthLargest {
private:
    int k;
    priority_queue<int, vector<int>, greater<int>> minHeap;  // min-heap
    
public:
    KthLargest(int k, vector<int>& nums) : k(k) {
        for (int num : nums) {
            add(num);  // reuse add() logic to build initial heap
        }
    }

    int add(int val) {
        minHeap.push(val);
        // Keep only the k largest elements
        if ((int)minHeap.size() > k) {
            minHeap.pop();  // remove the smallest
        }
        // Top of min-heap = kth largest
        return minHeap.top();
    }
};

/**
 * Your KthLargest object will be instantiated and called as such:
 * KthLargest* obj = new KthLargest(k, nums);
 * int param_1 = obj->add(val);
 */