class Solution {
public:
    vector<string> topKFrequent(vector<string>& words, int k) {
        unordered_map<string, int> freq;
        for (const string& w : words) {
            freq[w]++;
        }
        // Min-heap: keep k most frequent words
        // Min is at top → easily remove the least frequent when size > k
        // Comparator: min-heap by frequency; for ties → max by lex (reverse lex at top)
        auto cmp = [](const pair<int,string>& a, const pair<int,string>& b) {
            return a.first != b.first
                ? a.first > b.first // min-heap: lower freq at top
                : a.second < b.second; // for ties: larger lex at top (will be popped first)
        };
        priority_queue<pair<int,string>, vector<pair<int,string>>, decltype(cmp)> minHeap(cmp);
        for (auto& [word, count] : freq) {
            minHeap.push({count, word});
            if ((int)minHeap.size() > k) {
                minHeap.pop(); // remove least desirable element
            }
        }
        // Extract results in reverse order (min-heap gives smallest first)
        vector<string> result(k);
        for (int i = k - 1; i >= 0; i--) {
            result[i] = minHeap.top().second;
            minHeap.pop();
        }
        return result;
    }
};