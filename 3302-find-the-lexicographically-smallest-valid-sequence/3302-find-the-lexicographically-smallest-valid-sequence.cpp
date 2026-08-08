class Solution {
public:
    vector<int> validSequence(string word1, string word2) {
        int n = word1.size(), m = word2.size();
        // Step 1: Build suffMatch[i] = how many chars of word2's suffix
        // can be matched starting from word1[i], going right
        // suffMatch[i]: matching word2[m-suffMatch[i]..m-1] with word1[i..n-1]
        vector<int> suffMatch(n + 1, 0);
        int j = m - 1;
        for (int i = n - 1; i >= 0; i--) {
            suffMatch[i] = suffMatch[i + 1];
            if (j >= 0 && word1[i] == word2[j]) {
                suffMatch[i]++;
                j--;
            }
        }
        // Step 2: Greedy — match word2 prefix from left, use wild card optimally
        vector<int> result(m);
        bool usedWild = false;
        int wi = 0;  // current position in word2 (prefix matched so far)
        for (int i = 0; i < n && wi < m; i++) {
            if (word1[i] == word2[wi]) {
                // Exact match: take this index
                result[wi++] = i;
            } else if (!usedWild) {
                // Can we use wild card here?
                // After placing wild card at position wi (using index i),
                // we need suffMatch[i+1] >= m - wi - 1
                // (remaining word2 chars after wi: m - wi - 1)
                if (suffMatch[i + 1] >= m - wi - 1) {
                    // Use wild card at position wi, index i
                    result[wi++] = i;
                    usedWild = true;
                    // Now greedily fill remaining m-wi positions
                    // using suffix match from i+1 onward
                    for (int k = i + 1; k < n && wi < m; k++) {
                        if (word1[k] == word2[wi]) {
                            result[wi++] = k;
                        }
                    }
                    break;
                }
            }
        }
        // Check if we matched all of word2
        if (wi < m) return {};
        return result;
    }
};