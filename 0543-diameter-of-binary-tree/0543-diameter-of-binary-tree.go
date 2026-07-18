/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    best := 0
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        leftDepth  := dfs(node.Left)
        rightDepth := dfs(node.Right)
        // Diameter through current node = left edges + right edges
        if leftDepth+rightDepth > best {
            best = leftDepth + rightDepth
        }
        // Return depth of this subtree to parent
        return 1 + max(leftDepth, rightDepth)
    }
    dfs(root)
    return best
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}