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
    vector<vector<int>> levelOrder(TreeNode* root) {
        vector<vector<int>> result;
        if (root == nullptr) return result;
        queue<TreeNode*> bfsQueue;
        bfsQueue.push(root);
        while (!bfsQueue.empty()) {
            // Snapshot size so we only process nodes at the current level
            int layerSize = bfsQueue.size();
            vector<int> currentLevel;
            currentLevel.reserve(layerSize);
            for (int i = 0; i < layerSize; i++) {
                TreeNode* node = bfsQueue.front();
                bfsQueue.pop();
                currentLevel.push_back(node->val);
                if (node->left)  bfsQueue.push(node->left);
                if (node->right) bfsQueue.push(node->right);
            }
            result.push_back(move(currentLevel));
        }
        return result;
    }
};