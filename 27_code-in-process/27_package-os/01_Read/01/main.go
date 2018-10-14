package main

import "os"

func main() {
	src, err := os.Open("E:/GOPATH/src/GolangTraining-master2/27_code-in-process/27_package-os/01_Read/01/src.txt")
	if err != nil {
		panic(err)
	}

	defer src.Close()

	dst, err := os.Create("E:/GOPATH/src/GolangTraining-master2/27_code-in-process/27_package-os/01_Read/01/dst.txt")
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	bs := make([]byte, 5)
	src.Read(bs)
	dst.Write(bs)
}
