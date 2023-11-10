package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)

	go func() {
		defer fmt.Println("子进程结束")
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("子进程正在进行", i, len(c))
		}
		// 关闭chan
		close(c)
	}()

	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		m := <-c
		fmt.Println("m", m)
	}
	fmt.Println("主进程结束")
}
