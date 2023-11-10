package main

import "fmt"

func main() {
	// 无缓冲channel
	c := make(chan int)

	go func() {
		defer fmt.Println("groutine 结束")
		c <- 100

	}()
	num := <-c
	fmt.Println("num", num)
}
