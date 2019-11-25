package tree
//方法可以写在别的文件里，一个结构的方法可以分散在多个文件李
//树的中序遍历
func (node *Node)Traverse() {
	if node == nil {
		return
	}
	//java 或 c++要判断左子树或右子树不为空
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}