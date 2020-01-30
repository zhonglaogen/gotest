package main

import (
	"fmt"
	"time"
	"math/rand"
)

func generator() chan int {
	out := make(chan int)

	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		//判断是否无收据了,第二种方法用range
		//n, ok := <-c
		//if !ok{
		//	break
		//}
		time.Sleep(time.Second)
		fmt.Printf("worker %d is receive %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {

	//c1 c2 both nil
	var c1, c2 = generator(), generator()
	//w := createWorker(0)
	//var w chan int //nil chan nil的chan永远不会被select到
	var worker = createWorker(0)
	//非阻塞拿到chan的值，如果没有也不会一直等着拿值，而是走deafult
	//n := 0
	//此处主要作用是通过chan 的nil特性，只有拿到数据才执行 worker <- n
	//会发生的问题，如果消费者消费的太慢，会覆盖原来的值，导致消费的数量变少
	//解决办法，吧所有的value都存起来
	//hasValue := false
	var values []int
	//返回为一个Chan，时间结束后chan会收到值
	tm := time.After(10 * time.Second)
	//定时任务
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		//if hasValue{
		//	activeWorker = worker
		//}
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
			//hasValue = true
			//会阻塞住，直到等到可以发为止
			//w <- n
		case n := <-c2:
			values = append(values, n)
			//hasValue = true
			//w <- n
		case activeWorker <- activeValue:
			values = values[1:]
			//hasValue = false
		case <- time.After(800 * time.Millisecond):
			//两次select的时间差
			fmt.Println("timeout")
		case <- tick:
			fmt.Println("queue length = ",len(values))
		case <-tm:
			fmt.Println("bye")
			return


		}
	}
}
