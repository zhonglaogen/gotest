package basic

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

//返回错误 panic是抛出异常
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
	return 0, fmt.Errorf("unsupport %s"+op)
	}
	
}

//能返回两参数,可以定义别名
func div(a, b int) (q, r int) {
	return a / b, a % b
}

//智能返回参数
func div1(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

//函数式编程
func apply(op func(int, int) int, a, b int) int  {
	p :=reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("function name is %s args is (%d, %d)",opName, a, b)
	return op(a, b)
	
}
//重写Pow
func Pow(a, b int) int {
	return int(math.Pow(float64(a),float64(b)))
}

//默认参数：（x,y=3）python...   range 没有
//重载 没有
//没有操作符重载
//可选参数
//*args是可变参数，args接收的是一个tuple;
//
//**kw是关键字参数，kw接收的是一个dict
//操作符重载也没有

//可变参数列表

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
//go都是值传递，c即有值传递也有引用传递（int& a），java，python普遍引用传递，基本类型是值传递
//go配合指针可以做到引用传递(a *int),在传递对象时通常是传递对象的所指向数据的指针，所以封装类的时候也都是封装指针
//go的指针不可以想像c的数组一样相加减，所以简单
//值传递无效
func swap1(a, b int)  {
	b, a = a, b
}
//有效
func swap2(a, b *int)  {
	*b, *a = *a, *b
}

//有效
func swap3(a, b int) (int,int) {
	 return b, a
}


func main() {
	
	//fmt.Println(eval(1,2,"--"))
	div(13,3)
	fmt.Println(div(13, 2))

	//只用一个变量
	q, _:=div1(11, 8)
	fmt.Println(q)

	//3的4次方
	fmt.Println(apply(Pow,3,4))
	fmt.Println(apply(func(i int, i2 int) int {
		return i * i2
	},3,4))


	a, b := 3, 4
	//swap2(&a, &b)
	a, b = swap3(a, b)
	fmt.Println(a, b)
}
