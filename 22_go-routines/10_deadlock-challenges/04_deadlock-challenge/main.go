package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	//for {
	//	fmt.Println(<-c)
	//}
	for a := range c {
		fmt.Println(a)
	}
}
