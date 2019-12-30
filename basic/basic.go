package main

import (
	"fmt"
	"math/cmplx"
	"math"
	ioutil "io/ioutil"
	"strconv"
	os "os"
	"bufio"
)

//包的变量，不存在全局变量
var name = 1

var (
	aa = 1
	bb = 2
	cc = 3
	dd = 5
)

var a1, b1, v1 = 1, 2, 3

func variableZeroValue() {
	var a int
	var b string
	//必须有初始值,Printf可以跟变量格式
	//变量一旦定义必须用到
	fmt.Printf("%d %q", a, b)
}

func variableInitialValue() {
	//可以自动判定类别，类似py,py由上到下执行代码，全是kv,不同于js，js是贴签语言
	var a, b int = 1, 2
	var c string = "woaini"
	fmt.Println(a, b, c)
}

func variableTypeDeduction() {
	//混合定义
	var a, b, c, d, e = 1, 2, true, "aaa", 'w'
	fmt.Println(a, b, c, d, e)
}
func variableShorter() {
	//混合定义 := 是定义变量，在函数外面不能用这个定义 ：=， 必须用var
	a, b, c, d, e := 1, 2, true, "aaa", 'w'
	b = 5
	fmt.Println(a, b, c, d, e)
}

// bool string
//(u)int根据操作系统改变 （u）int8 （u）int32 （u）int64 uintptr
//byte rune(相当于char，在这里是32位)
//float32 float64 complex64 complex128 复数 实部加虚部

func euler() {
	//欧拉公式 e^i*@(角度)=cos@+isin@
	// |e^i*@|=1
	// e^i*pi=-1
	// e^0=1
	// e^i*@*pi*1/2=i
	// e^i*pi*3/2=-i
	//complex的实部和虚部都是浮点数字
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	//e的多少次方
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%.3f", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

const filename = "abc.text"

func consts() {
	//常量暂时不要大写
	const filename = "abc.text"
	//既可以做int 也可以float，文本替换,只有常量才可以
	const a, b = 3, 4
	math.Sqrt(a*a + b*b)

	const (
		aa = 3
		bb = 4
	)
}

func emums() {
	const (
		java       = 1
		python     = 2
		javascript = 3
		scala      = 4
	)
	//自增
	const (
		red    = iota
		yeoolw
		//跳值
		_
		blue
	)
	//b kb mb gb tb pb
	const (
		b  = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(java, python, java, scala)
	fmt.Println(red, yeoolw, blue)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func testif() {
	byte, error := ioutil.ReadFile("test.txt")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Printf("%s\n", byte)
	}

}

func testif2() {
	if byte, error := ioutil.ReadFile("test.txt"); error != nil {
		fmt.Println(error)
	} else {
		fmt.Printf("%s\n", byte)
	}
}

func testSwitch(score int) string {
	g := ""
	//switch可以没有表达式
	switch {
	case score < 0 || score > 100:
		//报错处理,后面代码不执行
		panic(fmt.Sprintf("Wrong is %s", score))
	case score < 60:
		g = "F"
	case score < 90:
		g = "B"

	}
	return g;
}

func testFor() int {
	sum := 0;
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return sum;
}

//求整数的 二进制
func convertToBinary(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		//数字转换为字符串
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func priintFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	bufio.NewScanner(file)

	scanner := bufio.NewScanner(file)

	//没有初始条件和递增条件，相当于while
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

//死循环
func forever() {
	//定义文本
	text := `dsa""fsa"a
     dsa
            fsaf
   asdg 

`
fmt.Println(text)
	for {
		fmt.Println("abc")
	}

}
func main() {
	//fmt.Println("hello word")
	//variableZeroValue()
	//variableInitialValue()
	//variableTypeDeduction()
	//variableShorter()

	euler()
	triangle()
	emums()

	testif2()

	//fmt.Println(testSwitch(1),testSwitch(-1))

	fmt.Println(testFor())
	fmt.Println(convertToBinary(13))

	priintFile("test.txt")

	//forever()
}
