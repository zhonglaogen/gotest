package main

import (
	"time"
	"fmt"
	"runtime"
)
//在命令行进入这个文件加 go run xx.go 运行程序 go  run -race xx.go 检测数据访问冲突
func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		//不传入i，就是闭包外面的i
		//传入i是因为 使用外部的i，会造成数组溢出，最后一次i ++ 的i会影响的所有线程

		go func(i int) {
			for  {
				a[i]++
				//强行让出资源，主线程结束，所有线程结束
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(10 * time.Millisecond)
	//主线程在读，其他线程在写
	fmt.Println(a)


}
