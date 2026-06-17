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
    bool isValidBST(TreeNode* root) {
        return validate(root, LLONG_MIN, LLONG_MAX);
    }
private:
    // Use long long bounds to safely handle INT_MIN/INT_MAX node values
    bool validate(TreeNode* node, long long lowerBound, long long upperBound) {
        // Empty subtree is trivially valid
        if (node == nullptr) return true;
        // Current node's value must strictly respect the inherited bounds
        if (node->val <= lowerBound || node->val >= upperBound) {
            return false;
        }
        // Left subtree: values must be < node->val (tighten upper bound)
        // Right subtree: values must be > node->val (tighten lower bound)
        return validate(node->left,  lowerBound, node->val)
            && validate(node->right, node->val,  upperBound);
    }
};