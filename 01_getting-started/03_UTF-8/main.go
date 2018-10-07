package main

import "fmt"

func main() {
	for i := 60; i < 122; i++ {
		//%q 对应的字符
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
