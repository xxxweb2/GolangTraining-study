package main

import (
	"io/ioutil"
	"log"
	"os"
	"fmt"
)

func main() {
	myStr := "hello Everybody"

	err := ioutil.WriteFile("E:/GOPATH/src/GolangTraining-master2/27_code-in-process/31_package-ioutil/02_WriteFile/xu2.txt", []byte(myStr), 0777)
	if err != nil {
		log.Fatalln("something went wrong!", err.Error())
	}
	f, err := os.Open("E:/GOPATH/src/GolangTraining-master2/27_code-in-process/31_package-ioutil/02_WriteFile/xu2.txt")
	if err != nil {
		log.Fatalln("couldn't read file!", err.Error())
	}
	defer f.Close()
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("Readall crashed!", err.Error())
	}
	fmt.Println(string(bs))
}
