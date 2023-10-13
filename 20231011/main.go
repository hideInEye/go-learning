package main

import (
	"fmt"
	"sync"
)

func main() {
	// 常量定义错误
	//var x = nil //error
	//var x  interface{}= nil 接口定义
	//var x map[string]string = nil map定义
	//var x []string = nil 切片定义
	//var x chan string = nil 管道定义
	//var x func() = nil // 函数定义

	//fmt.Println(x)

	// 切片跟数组不同
	//x := []string{"1"}
	//fmt.Println(x)
	//s := append(x, "2")
	//fmt.Println(s)
	//var x []string
	//fmt.Println(len(x))
	//s := append(x, "1")
	//fmt.Println(s)
	//h, w := 2, 4
	//
	//raw := make([]int, h*w)
	//for i := range raw {
	//	raw[i] = i
	//}
	////prints: [0 1 2 3 4 5 6 7] <ptr_addr_x>
	//table := make([][]int, h)
	//for i := range table {
	//	fmt.Println(raw[i*w : i*w+w])
	//	table[i] = raw[i*w : i*w+w]
	//}
	var wg sync.WaitGroup
	done := make(chan struct{}, 4)
	wq := make(chan interface{})
	workerCount := 2

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doit(i, wq, done, &wg)
	}

	for i := 0; i < workerCount; i++ {
		wq <- i
	}

	close(done)
	wg.Wait()
	fmt.Println("all done!")
}

func doit(workerId int, wq <-chan interface{}, done <-chan struct{}, wg *sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", workerId)
	defer wg.Done()
	for {
		select {
		case m := <-wq:
			fmt.Printf("[%v] m => %v\n", workerId, m)
		case <-done:
			fmt.Printf("[%v] is done\n", workerId)
			return
		}
	}
}
