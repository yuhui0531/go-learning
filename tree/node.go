package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node treeNode) print() {
	fmt.Print(node.value, " ")
}

//传值
func (node treeNode) setValue(value int) {
	node.value = value
}

// 传指针
func (node *treeNode) setValue2(value int) {
	if node == nil {
		fmt.Println("setting value to nil node. Ignored.")
		return
	}
	node.value = value
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

func main() {

	var root treeNode
	fmt.Println(root)

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}

	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	fmt.Println(root)
	root.left.right.setValue(4)
	root.print()
	fmt.Println()

	pRoot := &root
	pRoot.setValue2(100)
	pRoot.print()

	pRoot.setValue2(200)
	pRoot.print()

	var proot2 *treeNode
	proot2.setValue2(300)

	proot2 = &root
	proot2.setValue2(500)

	var root2 treeNode = treeNode{
		value: 0,
	}

	root2.left = &treeNode{
		value: 1,
	}

	root2.right = &treeNode{
		value: 2,
	}

	root2.left.right = &treeNode{
		value: 4,
	}

	root2.left.left = &treeNode{
		value: 3,
	}

	root2.right.left = &treeNode{
		value: 5,
	}

	root2.right.right = &treeNode{
		value: 6,
	}

	root2.traverse()

	/*
	   nodes := []treeNode{
	       {value: 3},
	       {},
	       {6, nil, &root},
	   }

	   fmt.Println(nodes)
	*/
}
