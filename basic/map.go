package basic

import "fmt"

func main() {
	//go的换行必须加入，不然就要在行尾部加入大括号
	m := map[string]string{
		"name":   "zlx",
		"age":    "11",
		"gender": "man",
	}
	fmt.Println("length",len(m))

	m2 := make(map[string]string) //m2 == empty map
	var m3 map[string]int //m3 == nil,nil可以和empty混用
	fmt.Println(m, m2, m3)

	for k, v := range m {
		//遍历是无序的，可能多次遍历的顺序也不一样
		fmt.Println(k,v)
	}
	fmt.Println("getting value")
	name, ok := m["name"]
	fmt.Println(name,ok)
	//如果不存在会输出空行（nil），用ok来判断是否有值,能两次重新定义ok是因为ok定义过，但是name1没定义过
	name1, ok := m["nam"]
	fmt.Println(name1,ok)

	fmt.Println("delete element")
	delete(m,"name")
	name, ok =m["name"]
	fmt.Println(name,ok)



}
