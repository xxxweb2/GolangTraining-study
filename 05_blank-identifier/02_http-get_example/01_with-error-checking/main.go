package main

import (
	"net/http"
	"github.com/gpmgo/gopm/modules/log"
	"io/ioutil"
	"fmt"
)

func main() {
	res, err := http.Get("http://www.baidu.com/")
	if err != nil {
		log.Fatal(err.Error())
	}
	page, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%s", page)
}
