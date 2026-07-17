/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
    result, path := [][]int{}, []int{}
    var dfs func(node *TreeNode, remaining int)
    dfs = func(node *TreeNode, remaining int) {
        if node == nil {
            return
        }
        // Add current node to path
        path = append(path, node.Val)
        remaining -= node.Val
        // Leaf check: if sum matches, record a copy of the current path
        if node.Left == nil && node.Right == nil && remaining == 0 {
            pathCopy := make([]int, len(path))
            copy(pathCopy, path)
            result = append(result, pathCopy)
        }
        // Recurse into children
        dfs(node.Left, remaining)
        dfs(node.Right, remaining)
        // Backtrack: remove current node before returning to parent
        path = path[:len(path)-1]
    }
    dfs(root, targetSum)
    return result
}