package piscine

func BTreeNodesCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + BTreeLevelCount(root.Left) + BTreeLevelCount(root.Right)
}

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(BTreeLevelCount(root.Left), BTreeLevelCount(root.Right))
}
