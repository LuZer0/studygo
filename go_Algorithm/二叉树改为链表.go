package main

func flatten(root *TreeNode)  {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	right := root.Right
	root.Right = root.Left
	root.Left = nil
	temp := root
	for temp.Right != nil {
		temp = temp.Right
	}
	temp.Right = right
}


