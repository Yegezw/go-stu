package main

import (
	"fmt"
	"time"
)

func test5() {
	fmt.Println("----------------------test4----------------------")

	demo1()
	fmt.Println()

	demo2()
	fmt.Println()

	demo3()
	fmt.Println()

	demo4()
	fmt.Println()
}

func demo1() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Println(time.Now())

	fmt.Println(timer.Stop())    // true, 因为 timer 还没有触发
	timer.Reset(2 * time.Second) // 重置 timer, 重新开始计时
	<-timer.C                    // 会等待 2 秒
	fmt.Println(timer.Stop())    // false, 因为 timer 已经触发了

	fmt.Println(time.Now())
}

func demo2() {
	// 1 秒后执行 func, 由新的 goroutine 执行
	timer := time.AfterFunc(
		time.Second,
		func() {
			fmt.Println(time.Now())
		},
	)

	fmt.Println(time.Now())
	time.Sleep(2 * time.Second) // 等待 2 秒, 确保定时器触发
	timer.Stop()
}

func demo3() {
	ch := make(chan struct{})

	go func() {
		time.Sleep(3 * time.Second)
		ch <- struct{}{} // 3 秒后向通道发送信号
	}()

	select {
	case <-ch:
		fmt.Println("Received signal")
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout!")
	}
}

func demo4() {
	ticker := time.NewTicker(time.Second)

	go func(ticker *time.Ticker) {
		for range ticker.C {
			fmt.Println(time.Now())
		}
	}(ticker)

	time.Sleep(5 * time.Second) // 等待 5 秒, 确保 ticker 触发多次
	ticker.Stop()               // 停止 ticker
}
