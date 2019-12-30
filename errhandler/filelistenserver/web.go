package main

import (
	"net/http"
	"zlx.com/laogen/gotest/errhandler/filelistenserver/filelisten"
	"os"
	"log"
)

type appHander func(writer http.ResponseWriter, request *http.Request) error
//包装，入参为目标函数，出参也为目标函数，还是执行原来的函数，但是增加了某些东西
func errWrapper(
	handler appHander) func(
		http.ResponseWriter, *http.Request) {

			return func(writer http.ResponseWriter, request *http.Request) {
				err :=  handler(writer, request)
				code := http.StatusOK
				if err != nil{
					log.Printf("Error handler request: %s",err)
					switch  {
					case os.IsNotExist(err):
						code = http.StatusNotFound
					default:
						code = http.StatusInternalServerError
					}
					http.Error(writer,
						http.StatusText(code),
						code)
				}
			}
}

//显示文件
func main() {
	http.HandleFunc("/list/",
		errWrapper(filelisten.HandlerFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
