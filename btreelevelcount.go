package piscine

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

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
	return 1 + intMax(BTreeLevelCount(root.Left), BTreeLevelCount(root.Right))
}
