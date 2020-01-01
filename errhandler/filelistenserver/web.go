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

		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}

		}()

		err := handler(writer, request)
		code := http.StatusOK
		if err != nil {
			log.Printf("Error handler request: %s", err)
			if userErr, ok := err.(userError); ok {
				http.Error(
					writer,
					userErr.Message(),
					http.StatusBadRequest,
				)
				return
			}
			
			switch {
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

//自定义错误类型
type userError interface {
	error
	Message() string
}


//显示文件
func main() {
	http.HandleFunc("/",
		errWrapper(filelisten.HandlerFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
