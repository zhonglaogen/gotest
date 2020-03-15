package main

import "fmt"

//切片就是视图
//不加长度就是切片
func updateSlices(s []int) {

	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:6]
	fmt.Println(s)
	fmt.Println(arr[2:])
	fmt.Println(arr[:7])
	fmt.Println(arr[:])
	fmt.Println("*******************************************")
	//改变视图的数据也会改变源数据，用切片不用传指针了，arr[:]传入的就不是数组就是切片了
	s1 := arr[2:]
	fmt.Println(s1)
	s2 := arr[:]
	fmt.Println(s2)

	updateSlices(s1)
	updateSlices(s2)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(arr)
	fmt.Println("*******************************************")

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)
	fmt.Println("*******************************************")

	fmt.Println("extending silce")
	arr[0], arr[2] = 0, 2
	s1 = arr[2:6]
	s2 = s1[3:5]
	//s1[4]直接取，取不出来，但是用slice可以取出来
	fmt.Println()
	fmt.Printf("arr=%v len(arr)=%d  cap(arr)=%d",arr,len(arr),cap(arr))
	fmt.Println()
	fmt.Printf("s1=%v len(s1)=%d  cap(s1)=%d",s1,len(s1),cap(s1))
	fmt.Println()
	fmt.Printf("s2=%v len(s2)=%d  cap(s2)=%d",s2,len(s2),cap(s2))
	fmt.Println()
	fmt.Println("*******************************************")

	//10顶替掉了7
	s3 := append(s2, 10)
	//超过了数组的cap，开辟新的arry，原来的数据会被拷过去
	//s4 和 s5 是一个不同的更大的数组，旧数组如果没人用就会被垃圾回收掉，s4和s5是对一个新数组的view

	//这个拷贝算上11一共增加了三个位置
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println(s3, s4, s5)
	fmt.Println(arr)
	//s2还是原数组的视图
	//fmt.Println(s2[:4])

	//新增数组的时候ptr,lenth,cap都会发生改变，由于会产生新的数组，ptr和cap，length都会变，所以必须有返回值
	//此次增加会覆盖上次s4的添加值
	fmt.Println(append(s4,23))
	fmt.Println(s4[:6])

}
