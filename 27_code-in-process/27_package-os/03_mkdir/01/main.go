package main

import (
	"os"
	"log"
)

func main() {
	err := os.Mkdir("E:/GOPATH/src/GolangTraining-master2/27_code-in-process/27_package-os/03_mkdir/01/test", 0x777)
	if err != nil {
		log.Fatalln("my program broke on mkdir: ", err.Error())
	}
	f, err := os.Create("E:/GOPATH/src/GolangTraining-master2/27_code-in-process/27_package-os/03_mkdir/01/test/test.txt")
	if err != nil {
		log.Fatalln("my program broke", err.Error())
	}
	defer f.Close()
	str := "put some phrase here."
	bs := []byte(str)
	_, err = f.Write(bs)
	if err != nil {
		log.Fatalln("error writing to file: ", err.Error())
	}
}
