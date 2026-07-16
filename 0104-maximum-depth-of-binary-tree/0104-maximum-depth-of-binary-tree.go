/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftDepth  := maxDepth(root.Left)
    rightDepth := maxDepth(root.Right)
    // Current node adds 1 level on top of the deeper subtree
    return 1 + max(leftDepth, rightDepth)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}