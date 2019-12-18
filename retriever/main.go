package main

import (
	"fmt"
	"zlx.com/laogen/gotest/retriever/mock"
	real2 "zlx.com/laogen/gotest/retriever/real"
)

//duck typing python是运行时检验，c++是编译时检验（template），java是继承
type AA interface {
	//这里不用加func关键字，因为这里本身就都是函数
	Get(url string) string
}
type BB interface {
	Post(url string,
		form map[string]string) string
}

const url  = "const url  = http://www.imooc.com"

func post(poster BB)  {
	poster.Post(url,
		map[string]string{
			"name": "ccmouse",
			"course": "golang",
		})
	
}
//接口的组合
type AABB interface {
	AA
	BB
	Conectcc(host string)
}
type CC interface {
	Conectcc(host string)
} 

func session(poster AABB)string  {
	poster.Post(url,map[string]string{
		"contents": "another faked  ... com",
	})
	return poster.Get(url)
}


func download(r AA) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r AA
	//r的肚子类，有一个类型，有一个值
	//相当于把等号后面的东西传入r中
	fmt.Printf("%T %v\n", r, r)
	r = mock.Retriever{"继承类"}
	fmt.Printf("%T %v\n", r, r)
	//接口通常不用指针，因为接口内部都含有一个指针
	r = &real2.Retriever{}
	fmt.Printf("%T %v\n", r, r)
	//fmt.Println(download(r))
	fmt.Println(download(mock.Retriever{"啦啦啦"}))

	//type assertion，通过.括号取出体内的类型
	realRetriever := r.(*real2.Retriever)
	fmt.Println(realRetriever.UserAgent)

	//防止报错，优雅的写
	if realRetriever2, ok := r.(mock.Retriever);ok{
		fmt.Println(realRetriever2)
	}else {
		fmt.Println("not a mock Retriever")
	}
	inspect(r)
	r = mock.Retriever{"测试重写string方法"}
	inspect(r)
	fmt.Println("try session")
	s := mock.Retriever{"AABB 组合测试"}
	fmt.Println(session(s))
}

func inspect(r AA) {
	fmt.Println("Inspecting ",r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Print(" > type switch")
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("false Retriever",v.Contents)
	case *real2.Retriever:
		fmt.Println("true Retriever",v.UserAgent)
	}
	fmt.Println()
}
