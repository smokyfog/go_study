package main

import "tree"

func main()  {
	var root tree.TreeNode
	root = TreeNode{ value: 3 }
	root.Left = &TreeNode{}
	root.Right = &TreeNode{5, nil, nil}
	root.Right.left = new(TreeNode)
	root.Left.right = createNode(2)
	// nodes := []TreeNode {
	// 	{ value: 3 },
	// 	{},
	// 	{6, nil, &root},
	// }
	// root.print()
	root.right.left.setValue(4)
	// root.right.left.print()
	// fmt.Println(nodes)

	// pRoot := &root
	// pRoot.print()
	// pRoot.setValue(200)
	// pRoot.print()

	// var pRoot *TreeNode
	// pRoot.setValue(200)
	// pRoot = &root
	// pRoot.setValue(300)
	// pRoot.print()

	root.traverse()
}