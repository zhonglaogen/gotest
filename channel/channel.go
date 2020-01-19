package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
		for n := range c{
			//判断是否无收据了,第二种方法用range
			//n, ok := <-c
			//if !ok{
			//	break
			//}
			fmt.Printf("worker %d is receive %c\n", id, n)
		}
}


//Channel是两个携程之间的数据传送，一个线程会死锁
//chan<- 表示告诉外面的人发数据，类似与java的泛型,但是这个Channel返回后从此以后就无法收数据
//<-chan 表示返回的chan只能收数据
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo()  {
	//这样的Channel默认是nil的，select会用到
	//var c chan int

	var channels [10]chan<- int
	for i := 0; i < 10; i++{
		channels[i] = createWorker(i)

	}

	//发完了，要是没有其他线程收，自己是收不到的，会死锁
	//c <- 1
	//c <- 2
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i

	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i

	}

	//另一个线程来不及收，就结束了
	time.Sleep(time.Millisecond)
}

//利用缓冲区
func bufferChannel() {
	//第二个参数是缓冲区，能发三条数据，但是第四个就会报错,缓提区提高性能
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)

	//如果当g'前线程只发数据，没有其他线程收数据，那么这个线程就会死掉
//	发完数据后，虽然是携程，也需要切换，一直等人来收，也很耗费资源
}

//永远是发送发来close 的
func closeChannel()  {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	//告诉接收方不发数据了，但是接收方还是可以收到数据的，收到是方框（字符），数字就是0
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	//csp模型，并发模型，不用通过共享内存来通信，要通过通信来共享内存

	//传递Channel
	//chanDemo()
	//缓冲区Channel
	//bufferChannel()
	//Channelclose，用range收
	closeChannel()
}
