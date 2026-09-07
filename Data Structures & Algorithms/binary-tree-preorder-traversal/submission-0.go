/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func preorderTraversal(root *TreeNode) []int {
    res := []int{}

	var preorder func(node *TreeNode)
	preorder= func (node* TreeNode){
		if node==nil{
			return
		}
		// Process first
		res= append(res,node.Val)

		// Check Left
		preorder(node.Left)
		// Check Right
		preorder(node.Right)
	}
	preorder(root)
	return res
}
