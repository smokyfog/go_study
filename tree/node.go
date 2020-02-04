package tree

import "fmt"

type TreeNode struct {
	Value int
	Left, Right *TreeNode
	Left, Right *TreeNode
}

// 为结构定义方法
func (node TreeNode) Print() {
	fmt.Println(node.value)
}

func (node *TreeNode) SetValue(value int) {
	if node == nil {
		fmt.Println("setting value to nil node...")
		return
	}
	node.value = value
	fmt.Println("node.value", node.value)
}

func CreateNode(value int) *TreeNode {
	return &TreeNode{value: value}
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}
