/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	var traverseLeftFirst func(node *TreeNode)
	traverseLeftFirst = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverseLeftFirst(node.Left)
		res = append(res, node.Val)
		traverseLeftFirst(node.Right)
	}
	traverseLeftFirst(root)
	return res
}
