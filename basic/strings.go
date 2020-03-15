package main

import (
	"fmt"
	"unicode/utf8"
)

//unicode码是通用字符表示码，utf-8是一种编解码方式，英文一个字节，中文三个字节，rune就是int32的别名


func main() {
	s := "YES我是钟老根！"
	fmt.Println(len(s))
	fmt.Printf("%s",[]byte(s))
	fmt.Println()
	fmt.Printf("%X",[]byte(s))
	fmt.Println()
	fmt.Println("************1***************")
	for _, b := range []byte(s){
		fmt.Printf("%X ",b)
	}
	fmt.Println()
	fmt.Println("************2***************")
	for i, ch := range s { //ch is rune
		fmt.Printf("%d-%X ", i, ch)
	}
	fmt.Println()
	fmt.Println("************3***************")

	fmt.Println(utf8.RuneCountInString(s))

	fmt.Println()
	fmt.Println("*************4**************")

	bytes := []byte(s)
	for len(bytes) > 0{
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Println(ch)
		fmt.Printf("%c ", ch)
		fmt.Println()
	}
	fmt.Println("*************5**************")
	//另外开了一个数组，村这些字符
	for _, ch := range ([]rune(s)){
		fmt.Printf("%c", ch)
	}


}
