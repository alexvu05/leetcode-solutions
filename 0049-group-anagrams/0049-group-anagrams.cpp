class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        // Map from sorted-key → list of original strings sharing that key
        unordered_map<string, vector<string>> groups;
        for (const string& s : strs) {
            // Sorting the characters produces a canonical key for all anagrams
            string key = s;
            sort(key.begin(), key.end());
            groups[key].push_back(s);
        }
        // Flatten the map values into the result
        vector<vector<string>> result;
        result.reserve(groups.size());
        for (auto& [key, group] : groups) {
            result.push_back(move(group));
        }
        return result;
    }
};