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
    vector<double> averageOfLevels(TreeNode* root) {
        vector<double> result;
        if (root == nullptr) return result;
        queue<TreeNode*> bfsQueue;
        bfsQueue.push(root);
        while (!bfsQueue.empty()) {
            int layerSize = bfsQueue.size();
            double sum = 0;  // use double to avoid overflow accumulation
            for (int i = 0; i < layerSize; i++) {
                TreeNode* node = bfsQueue.front();
                bfsQueue.pop();
                sum += node->val;
                if (node->left)  bfsQueue.push(node->left);
                if (node->right) bfsQueue.push(node->right);
            }
            result.push_back(sum / layerSize);
        }
        return result;
    }
};