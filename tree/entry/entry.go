package main

import (
	"fmt"
	"zlx.com/laogen/gotest/tree"
)
//一个目录下面只能有一个包，但是名字可以和目录不一样
//方法名,类型,属性，首字母大写代表public，小写为private

//扩充，继承（定义别名，使用）
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode)postOrder()  {

	if myNode == nil  || myNode.node == nil{
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()

}

func main() {
	var root tree.Node
	fmt.Println(root)
	root = tree.Node{Value:3}
	root.Left = &tree.Node{2,nil,nil}
	root.Right = &tree.Node{}
	//用.来访问，如果是c++ 指针指向的实例要用->
	root.Right.Left = new(tree.Node)
	root.Left.Right=tree.CreateNode(2)

	//这个是把地址复制过去
	root.SetValue(1)
	root.Print()
	fmt.Println()

	//nil指针可以传进去
	var pRoot *tree.Node
	pRoot.SetValue(100)
	pRoot = &root
	pRoot.SetValue(300)
	fmt.Println("zhong")
	root.Traverse()

	fmt.Println("hou")
	mynode := myTreeNode{&root}
	mynode.postOrder()


	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println(nodeCount)
}