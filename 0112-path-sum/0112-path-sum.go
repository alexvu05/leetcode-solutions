/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    remaining := targetSum - root.Val
    // Reached a leaf: check if remaining sum is exactly 0
    if root.Left == nil && root.Right == nil {
        return remaining == 0
    }
    // Recurse into non-nil children
    return hasPathSum(root.Left, remaining) || hasPathSum(root.Right, remaining)
}