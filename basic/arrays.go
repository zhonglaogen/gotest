package main

import "fmt"

//数组是值传递,数组是值类型的，[10]int 和 [20]int 是不同的类型
func change(arry [5]int) {
	arry[0] = 20
}

//这样传的是指针，可以改变
func change3(arry *[5]int) {
	arry[0] = 20
}

//切片为引用传递
func change2(arry []int) {
	arry[0] = 20
}

//定长数组，变长切片
func main() {
	var arry1 = []int{1, 2, 3}
	var arry [5]int
	//这样必须赋值
	arry2 := [5]int{2, 3, 4, 5, 6}
	//编译器数几个值
	arry3 := [...]int{2, 3, 4, 5, 7}

	var grid [2][3]int
	fmt.Println(arry, arry2, arry3, grid)

	for i := 0; i < len(arry3); i++ {
		fmt.Println(arry3[i])
	}

	for i := range arry3 {
		fmt.Println(arry3[i])
	}
	for i, v := range arry3 {
		fmt.Println(i, v)
	}

	//下划线表示不用的值
	for _, v := range arry3 {
		fmt.Println(v)
	}

	change(arry2)
	fmt.Println(arry2[0])
	change2(arry1)
	fmt.Println(arry1[0])

}
