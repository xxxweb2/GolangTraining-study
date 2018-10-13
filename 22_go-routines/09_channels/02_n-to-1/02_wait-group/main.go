package main

import (
	"sync"
)
var wg sync.WaitGroup

var c = make(chan int)
func main() {

	wg.Add(2)
	go a()
	go b()
	//var wg sync.WaitGroup
	//wg.Add(2)

	//go func() {
	//	for i := 0; i < 10; i++ {
	//		c <- i
	//	}
	//	wg.Done()
	//}()
	//
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		c <- i
	//	}
	//	wg.Done()
	//}()
	go func() {
	wg.Wait()
	//close(c)
	}()
	//for n := range c {
	//	fmt.Println(n)
	//}

}
func a() {
	for i := 0; i < 10; i++ {
		c <- i
	}
	wg.Done()
}

func b() {
	for i := 0; i < 10; i++ {
		c <- i
	}
	wg.Done()
}
