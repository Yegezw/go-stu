package main

import (
	"fmt"
	"time"
)

type rChannel <-chan int // 只读信道
type wChannel chan<- int // 只写信道

func test1() {
	fmt.Println("----------------------test1----------------------")

	// 求和
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := make(chan int, 2)

	go sum(arr[:len(arr)/2], res) // 前半部分
	go sum(arr[len(arr)/2:], res) // 后半部分

	x, y := <-res, <-res
	fmt.Printf("%d + %d = %d\n", x, y, x+y)

	// 10 个数的斐波那契数列
	c := make(chan int, 10)
	go fib1(cap(c), c)
	// 不断从信道接收值, 直到它被关闭
	for i := range c {
		fmt.Println(i)
	}
	// 仍然可以从 "关闭的管道" 中接收数据 (永远是零值)
	i, ok := <-c
	fmt.Println(ok, i)

	// 不以共享内存来通信, 而以通信来共享内存
	ch := make(chan int, 10)
	rc := rChannel(ch)
	wc := wChannel(ch)
	go func() {
		for i := 0; i < 10; i++ {
			wc <- i
		}
		close(ch)
	}()
	go func() {
		for i := range rc {
			fmt.Println(i)
		}
	}()
	time.Sleep(10 * time.Millisecond)

	// channel 又分为两类: 有缓冲 channel 和无缓冲 channel
	// 为了协程安全, channel 内部都会有一把锁来控制并发访问, channel 底层用队列来存储数据
	// 在一端关闭 channel 的时候, 该 channel 读端的所有 qoRoutine 都会收到 channel 已关闭的消息
	// [1] 无缓冲 channel 可以理解为同步模式: 即写入一个, 如果没有消费者在消费, 写入就会阻塞
	// [2] 有缓冲 channel 可以理解为异步模式: 即写入消息之后, 即使还没被消费, 只要队列没满, 就可继续写入
	num := 0
	lock := make(chan bool, 1) // 用 channel 模拟锁
	for i := 0; i < 1000; i++ {
		go add(lock, &num)
	}
	time.Sleep(10 * time.Millisecond)
	fmt.Println("num:", num) // 1000
}

func sum(slice []int, res chan int) {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	res <- sum
}

func fib1(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// 只应由发送者关闭信道, 而不应由接收者关闭
	// 向一个已经关闭的信道发送数据会引发程序 panic
	close(c) // 关闭信道
}

func add(lock chan bool, num *int) {
	lock <- true
	*num += 1
	<-lock
}
