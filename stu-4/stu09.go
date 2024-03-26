package main

import (
	"fmt"
	"time"
)

func test10() {
	fmt.Println("----------------------test10----------------------")

	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// ----------------------------------------------------------

type MyError struct {
	When time.Time
	What string
}

// 通常函数会返回一个 error 值
// 调用的它的代码应当判断这个错误是否等于 nil 来进行错误处理
// error 为 nil 时表示成功、非 nil 的 error 表示失败

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
