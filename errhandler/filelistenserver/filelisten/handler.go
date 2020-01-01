package filelisten

import (
	"net/http"
	"os"
	"io/ioutil"
	"strings"
)

type userError string

func (e userError) Error()string  {
	return e.Message()
}
func (e userError) Message()string  {
	return string(e)
}


//将业务独立的抽取出来
const prefix  = "/list/"
func HandlerFileList(writer http.ResponseWriter, request *http.Request)error {
	if strings.Index(
		request.URL.Path,prefix) != 0{
		return userError("path must start" + "with" + prefix)
	}
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		//panic(err)
		//http.Error(writer,
		//	err.Error(),
		//	http.StatusInternalServerError)
		//return
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil{
		//panic(err)
		return err
	}
	writer.Write(all)
	return nil
}
