package basic

import "fmt"

func printSlice(s []int)  {
	//2倍扩充
	fmt.Println(s,len(s),cap(s))
}

func main() {
	fmt.Println("creating sclice")
	//Zero value for sclice is nil
	var s []int

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2 * i + 1)
	}
	fmt.Println(s)

	//将数组指派给一个slice
	s1 := []int{1, 2, 3, 5}
	printSlice(s1)

	//创建指定大小的切片
	s2 := make([]int,16)
	printSlice(s2)

	s3 := make([]int, 10, 32)
	printSlice(s3)

	//扩容就是*2
	s4 := make([]int,3)
	printSlice(s4)
	printSlice(append(s4,1))


	fmt.Println("copy slice")
	//把s1的值copy给s1
	copy(s2,s1)
	printSlice(s2)

	fmt.Println("deleting element from slice")

	//删除位置i的元素，后面的元素往前移动 slice[:i] + slice[4:] ...代表当前切片的所有元素
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println(front)
	printSlice(s2)

	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)





}
