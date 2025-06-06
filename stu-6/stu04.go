package main

import (
	"context"
	"fmt"
	"time"
)

/*
解决 goroutine 之间退出通知、元数据传递的功能
[1] 有上下级关系的 chan struct{} - 向下取消
[2] 有上下级关系的 key, val any  - 向上取值
type Context interface {
    // 截止时间 + 是否设置了截止
	Deadline() (deadline time.Time, ok bool)

    // 是否取消了
	Done() <-chan struct{}

    // 取消原因 Canceled / DeadlineExceeded
	Err() error

    // 获取值, 通过 WithValue 写入
	Value(key any) any
}
*/

func test4() {
	fmt.Println("----------------------test4----------------------")

	chancelDemo()
	deadlineDemo()
	timeoutDemo()
	withValueDemo()
}

// ----------------------------------------------------------

func chancelDemo() {
	ctx, cancel := context.WithCancel(context.Background())
	go Watch(ctx, "watch1")
	go Watch(ctx, "watch2")

	time.Sleep(2 * time.Millisecond)
	cancel() // 取消上下文, 停止所有 goroutine
}

func deadlineDemo() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Millisecond))
	defer func() {
		// 等待 watch3 + watch4 执行完
		time.Sleep(10 * time.Millisecond)
		cancel() // 也可以手动取消上下文, 停止所有 goroutine
	}()

	go Watch(ctx, "watch3")
	go Watch(ctx, "watch4")
}

func timeoutDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Millisecond)
	defer func() {
		// 等待 watch5 + watch6 执行完
		time.Sleep(10 * time.Millisecond)
		cancel() // 也可以手动取消上下文, 停止所有 goroutine
	}()

	go Watch(ctx, "watch5")
	go Watch(ctx, "watch6")
}

func Watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "is stopped")
			return
		default:
			fmt.Println(name, "is running")
		}
	}
}

// ----------------------------------------------------------

func withValueDemo() {
	ctx := context.WithValue(context.Background(), "name", "zs")
	go subRun(ctx)
	time.Sleep(2 * time.Millisecond)
}

func subRun(ctx context.Context) {
	fmt.Printf("name is %s\n", ctx.Value("name").(string))
}
