package main

import "io/ioutil"

func main() {
	err := ioutil.WriteFile("E:/GOPATH/src/GolangTraining-master2/27_code-in-process/31_package-ioutil/02_WriteFile/xu.txt", []byte("hello world"),
	0777)
	if err != nil {
		panic("something went wrong")
	}

}
