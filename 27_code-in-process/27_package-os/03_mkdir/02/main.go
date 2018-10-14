package main

import (
	"os"
	"log"
)

func main() {
	err := os.Mkdir("E:/GOPATH/src/GolangTraining-master2/27_code-in-process/27_package-os/03_mkdir/01/test2", os.ModePerm)
	if err != nil {
		log.Fatalln("my program broke on mkdir: ", err.Error())
	}

}
