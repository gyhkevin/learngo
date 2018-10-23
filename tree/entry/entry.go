package main

import (
	"fmt"
	"imooc.com/kevin/learngo/tree"
	"golang.org/x/tools/container/intsets"
)

type MyTreeNode struct {
	node *tree.Node
}

func (myNode *MyTreeNode) postOrder()  {
	if myNode == nil || myNode.node == nil{
		return
	}

	left := MyTreeNode{myNode.node.Left}
	right := MyTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func testSparse() {
	s := intsets.Sparse{}
	s.Insert(1)
	s.Insert(10000)
	s.Insert(100000000)
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(1000000))
}

func main()  {
	var root tree.Node

	root = tree.Node{Value:3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	nodes := []tree.Node {
		{Value:3},
		{},
		{6,nil,&root},
	}
	fmt.Println(nodes)
	root.Print()
	fmt.Println()

	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()

	//root.Print()
	//root.SetValue(200)
	//pRoot := &root
	//pRoot.Print()
	//pRoot.SetValue(300)
	/*var pRoot *Node
	pRoot.SetValue(200)
	pRoot = &root
	pRoot.SetValue(300)
	pRoot.Print()*/

	root.Traverse()
	fmt.Println()
	myRoot := MyTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
}