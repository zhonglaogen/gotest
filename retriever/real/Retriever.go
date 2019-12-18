package real

import (
	"time"
	"net/http"
	"net/http/httputil"
)

type Retriever struct {
	UserAgent string
	//毫秒
	TimeOut time.Duration
}

func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err :=  httputil.DumpResponse(resp, true)
	//读完要关掉
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(result)
}

