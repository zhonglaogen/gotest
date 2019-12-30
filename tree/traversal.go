package tree

import "fmt"

//方法可以写在别的文件里，一个结构的方法可以分散在多个文件李
//树的中序遍历
func (node *Node)Traverse() {
	node.TraverseFunc(func(node *Node) {
		node.Print()
	})
	fmt.Println()
}
//函数也可以继承，遍历的过程中做任何事情
func (node *Node) TraverseFunc(f func(*Node))  {
	if node == nil {
		return
	}
	//java 或 c++要判断左子树或右子树不为空
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}