package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	res, _ := http.Get("http://www.baidu.com/")
	page, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("%s", page)
}
