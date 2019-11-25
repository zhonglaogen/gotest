package tree

import (
	"fmt"
)
//go既可以定义值接受也可以指针，方法调用不用改

type Node struct {
	Value int
	Left, Right *Node
}
//(node Node)相当于this指针，为接受者，值传递，所有都是值传递,
// 其实传的是指针，编译器会自动将指针指向的数据解析出来
func (node Node) Print()  {
	fmt.Print(node.Value," ")
}
func (node *Node) SetValue(value int)  {
	if node == nil {
		fmt.Println("nil ^^^^^^^^")
		return
	}
	node.Value=value
}



//go分配变量不必关心是栈还是堆，如果外面没人用就是栈，有人用就是堆
func CreateNode(value int) *Node {
	//若果是c++，这是局部变量地址，但是go可以用
	return &Node{Value:value}
}


