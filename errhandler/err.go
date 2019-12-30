package main

import (
	"fmt"
)

func err(){
	for i :=0; i < 100; i++  {
		//defer会打印出0-10，而不是10个10
		defer fmt.Println(i)
		if i==10 {
			fmt.Println("a")
			//这个异常随机时间抛出，结果表现在defer时的某一个时刻抛出来
			panic("dddd./")
		}
	}
}



func main() {
	err()
}
