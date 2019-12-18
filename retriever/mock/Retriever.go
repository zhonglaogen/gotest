package mock

import "fmt"

type Retriever struct {
	Contents string

}

func (r Retriever) String() string {
	return fmt.Sprintf(
		"重写stirng后:{Contents=%s}",r.Contents,
	)
}

func (r Retriever) Post(url string,
	form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

func (r Retriever) Conectcc(host string) {
	fmt.Println("pass")
}

//实现的接口及时不同名，只要相同方法也可以
func (r Retriever) Get(url string) string {
	return r.Contents
}



