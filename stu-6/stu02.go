package main

import (
	"fmt"
	"time"
)

func test2() {
	fmt.Println("----------------------test2----------------------")

	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // 读 c
		}
		quit <- 0
	}()

	fib2(c, quit)
}

func fib2(c, quit chan int) {
	x, y := 0, 1
	for {
		// select 使 Go 程可以等待多个通信操作
		// 阻塞到某个分支可以继续执行为止, 当多个分支都准备好时会随机选择一个执行
		// 当 select 中的其它分支都没有准备好时, default 分支就会执行
		select {
		case c <- x: // 写 c
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("wait 10 ms")
			time.Sleep(10 * time.Millisecond)
		}
	}
}
