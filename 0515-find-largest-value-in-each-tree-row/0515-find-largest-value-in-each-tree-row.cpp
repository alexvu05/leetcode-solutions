/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    vector<int> largestValues(TreeNode* root) {
        vector<int> result;
        if (root == nullptr) return result;
        queue<TreeNode*> bfsQueue;
        bfsQueue.push(root);
        while (!bfsQueue.empty()) {
            int layerSize = bfsQueue.size();
            int maxVal    = INT_MIN;  // track max for current level
            for (int i = 0; i < layerSize; i++) {
                TreeNode* node = bfsQueue.front();
                bfsQueue.pop();
                // Update max for this level
                maxVal = max(maxVal, node->val);
                if (node->left)  bfsQueue.push(node->left);
                if (node->right) bfsQueue.push(node->right);
            }
            result.push_back(maxVal);
        }
        return result;
    }
};