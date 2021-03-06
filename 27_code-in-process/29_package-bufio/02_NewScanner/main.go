package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
)

func main() {
	src, err := os.Open("E:/GOPATH/src/GolangTraining-master2/test/1.txt")
	if err != nil {
		log.Printf("error opening source file: %v", err)
	}
	defer src.Close()

	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(">>>", line)
	}

}
