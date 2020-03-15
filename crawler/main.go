package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	resp, err := http.Get(
		"http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//输出所有响应，包括头
	//httputil.DumpResponse()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code",
			resp.StatusCode)
		return

	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", all)

}
