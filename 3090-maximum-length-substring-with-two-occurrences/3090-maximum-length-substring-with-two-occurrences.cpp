class Solution {
public:
    int maximumLengthSubstring(string s) {
        array<int, 26> freq{};
        int left = 0;
        int best = 0;
        for (int right = 0; right < (int)s.size(); right++) {
            freq[s[right] - 'a']++;
            // Shrink window until s[right]'s frequency is at most 2
            while (freq[s[right] - 'a'] > 2) {
                freq[s[left] - 'a']--;
                left++;
            }
            best = max(best, right - left + 1);
        }
        return best;
    }
};