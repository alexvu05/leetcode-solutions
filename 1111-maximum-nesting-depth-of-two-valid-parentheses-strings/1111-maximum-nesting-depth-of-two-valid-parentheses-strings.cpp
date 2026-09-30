class Solution {
public:
    vector<int> maxDepthAfterSplit(string seq) {
        int n = seq.size();
        vector<int> answer(n);
        int depth = 0;

        for (int i = 0; i < n; i++) {
            if (seq[i] == '(') {
                depth++;
                // Assign to group based on parity of current depth
                // Odd depth → group 0, even depth → group 1
                answer[i] = depth % 2;
            } else {
                // ')': assign to same group as matching '('
                // Current depth before decrement corresponds to matching '('
                answer[i] = depth % 2;
                depth--;
            }
        }

        return answer;
    }
};