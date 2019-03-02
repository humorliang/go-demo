package main

import (
	"net/http"
	"bytes"
	"fmt"
)

/*
Convert io.ReadCloser to a String
将io流变为string
*/
func main() {
	resp, _ := http.Get("http://www.baidu.com")
	//创建buf
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	fmt.Printf(buf.String())
}
