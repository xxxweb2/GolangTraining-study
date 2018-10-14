package main

import (
	"os"
	"log"
)

func main() {
	dst, err := os.Create("E:/GOPATH/src/GolangTraining-master/27_code-in-process/27_package-os/02_Write/02/hello.txt")
	if err != nil {
		log.Fatalln("error creating destination file: ", err.Error())
	}
	defer dst.Close()

	bs := []byte("put some phrase here.")

	_, err = dst.Write(bs)
	if err != nil {
		log.Fatalln("error writing to file: ", err.Error())
	}
}
