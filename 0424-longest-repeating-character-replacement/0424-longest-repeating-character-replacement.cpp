class Solution {
public:
    int characterReplacement(string s, int k) {
        array<int, 26> freq{};  // frequency of each character in the window
        int left = 0;
        int maxFreq = 0;        // highest frequency of any single char in window
        int best = 0;
        for (int right = 0; right < (int)s.size(); right++) {
            freq[s[right] - 'A']++;
            maxFreq = max(maxFreq, freq[s[right] - 'A']);
            // Characters to replace = window_size - maxFreq
            // Shrink window if replacements needed exceed k
            int windowSize = right - left + 1;
            if (windowSize - maxFreq > k) {
                freq[s[left] - 'A']--;
                left++;
            }
            best = max(best, right - left + 1);
        }
        return best;
    }
};