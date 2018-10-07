package main

import "fmt"

func main() {
	x := 0

	//匿名函数
	increment := func() int {
		x ++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
