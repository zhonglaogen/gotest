package main

import (
	"fmt"
	"time"
	"sync"
)

type atomic struct {
	value int
	lock sync.Mutex
}

func (ato *atomic) increment()  {
	//这样做，保证只锁了部分代码，defer在函数调用结束就执行
	func(){
		ato.lock.Lock()
		defer ato.lock.Unlock()
		ato.value++
	}()


}
func (ato *atomic) get()int  {
	ato.lock.Lock()
	defer ato.lock.Unlock()
	return int(ato.value)
}

func main() {
	var a atomic
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	//别人在写，我在读，有可能发生别人在写的时候我要读
	//枷锁后执行 go run -race aotomic.go 就没有冲突了
	fmt.Println(a.get())
}
