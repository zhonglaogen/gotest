package main

import (
	"fmt"
	"errors"
)

func tryRecover()  {

	defer func() {
		r := recover()
		if err, ok := r.(error); ok{
			fmt.Println("Error occured:", err)
		}else if err == nil{
			fmt.Println("无异常")
		} else {
			fmt.Println("***")
			panic(r)
		}
	}()
	panic(errors.New("this is a error"))
}

func main() {
	tryRecover()
	
}
