package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, w worker) {
	for n := range w.in {
		//判断是否无收据了,第二种方法用range
		//n, ok := <-c
		//if !ok{
		//	break
		//}
		fmt.Printf("worker %d is receive %c\n", id, n)
		//第一次大会卡在这里没人收,此线程是阻塞的,第二次就没法收到了,死锁
		//go func() {done <- true}()
		w.done()

	}
}

//解决需要睡眠的问题

type worker struct {
	in   chan int
	//wg *sync.WaitGroup
	done func()
}

//Channel是两个携程之间的数据传送，一个线程会死锁
//chan<- 表示告诉外面的人发数据，类似与java的泛型,但是这个Channel返回后从此以后就无法收数据
//<-chan 表示返回的chan只能收数据
func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in:   make(chan int),
		done : wg.Done,
	}
	//每个线程只能拥有把一条Channel
	go doWorker(id, w)
	return w
}

func chanDemo() {
	//这样的Channel默认是nil的，select会用到
	//var c chan int
	var wg sync.WaitGroup
	//add是有多少个任务
	wg.Add(20)

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i,&wg)

	}


	//发完了，要是没有其他线程收，自己是收不到的，会死锁
	//c <- 1
	//c <- 2
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		//	不在这里收是为了不让他们顺序打印
	}


	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i

	}
	wg.Wait()

	//	统一发送完数据后，在这里收
	//for _, worker := range workers {
	//	<-worker.done
	//	<-worker.done
	//
	//}

}

func main() {
	//csp模型，并发模型，不用通过共享内存来通信，要通过通信来共享内存


	//传递Channel
	chanDemo()
}
