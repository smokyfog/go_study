package main

import "fmt"

type treeNode struct {
	value int
	left, right *treeNode
}

// 为结构定义方法
func (node treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("setting value to nil node...")
		return
	}
	node.value = value
	fmt.Println("node.value", node.value)
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main()  {
	var root treeNode
	root = treeNode{ value: 3 }
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	// nodes := []treeNode {
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

	// var pRoot *treeNode
	// pRoot.setValue(200)
	// pRoot = &root
	// pRoot.setValue(300)
	// pRoot.print()

	root.traverse()
}