package main

import (
	"net/http"
	httputil "net/http/httputil"
	"fmt"
)

func main() {

	request, err := http.NewRequest(http.MethodGet,
		"http://www.imooc.com", nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	client := http.Client{
		//重定向会执行的方法
		CheckRedirect: func(
			req *http.Request,
			via []*http.Request) error {

				fmt.Println("REQUEST:", req)
				//让他重定向返回nil
				return nil
		},
	}
	resp, err := client.Do(request)
	//resp, err := http.DefaultClient.Do(request)
	//非特定的浏览请求
	//resp, err :=  http.Get("http://www.imooc.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//知道responese的内容
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n",s)

}
