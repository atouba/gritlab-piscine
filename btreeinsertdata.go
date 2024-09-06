package piscine

type TreeNode struct {
	Parent, Right, Left *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	for curr := root; curr != nil; {
		if Compare(data, curr.Data) > 0 {
			if curr.Right != nil {
				curr = curr.Right
			} else {
				curr.Right = &TreeNode{curr, nil, nil, data}
				return curr.Right
			}
		} else {
			if curr.Left != nil {
				curr = curr.Left
			} else {
				curr.Left = &TreeNode{curr, nil, nil, data}
				return curr.Left
			}
		}
	}

	root = &TreeNode{nil, nil, nil, data}
	return root
}
