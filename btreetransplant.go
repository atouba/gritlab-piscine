package piscine

// func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
// 	*node = *rplc
// 	return root
// }

func isNodeLeft(node *TreeNode) bool {
	return node == node.Parent.Left
}

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == node {
		*root = *rplc
	} else if isNodeLeft(node) {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}
	return root
}
