package piscine

// Why does this work? Shouldn't it return the first node,
// since it keeps calling back...
// func BTreeMax(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	if BTreeMax(root.Right) == nil {
// 		return root
// 	}
// 	return BTreeMax(root.Right)
// 	// return nil
// }

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	curr := root
	for ; curr.Right != nil; curr = curr.Right {
	}
	return curr
}
