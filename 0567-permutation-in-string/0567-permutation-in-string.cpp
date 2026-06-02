class Solution {
public:
    bool checkInclusion(string s1, string s2) {
        int n1 = s1.length();
        int n2 = s2.length();
        if (n1 > n2) return false;
        std::vector<int> count1(26, 0);
        std::vector<int> count2(26, 0);
        for (int i = 0; i < n1; ++i) {
            count1[s1[i] - 'a']++;
            count2[s2[i] - 'a']++;
        }
        if (count1 == count2) return true;
        for (int i = n1; i < n2; ++i) {
            count2[s2[i] - 'a']++;
            count2[s2[i - n1] - 'a']--;
            if (count1 == count2) return true;
        }
        return false;
    }
};