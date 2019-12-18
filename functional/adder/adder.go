package main

import "fmt"

//v是局部变量，但是sum是自由变量，内部函数的sum指向外部，
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

//结果返回两个，一个是计算的结果值，一个返回的是带参数的函数
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder  {

	return func(v int) (int, iAdder) {
		return base +v, adder2(base)
	}
}

func main() {

	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + .. + %d = %d\n", i, a(i))
	}
}
