package main

import (
	"errors"
	"fmt"
	"time"
)

func test10() {
	fmt.Println("----------------------test10----------------------")

	if err := run(); err != nil {
		fmt.Printf("%T = %v\n", err, err)
		// *main.MyError = at 2025-06-05 18:05:06.042737 +0800 CST m=+0.000279751, it didn't work
	}

	e1 := errors.New("this is error1")
	e2 := fmt.Errorf("this is %s", "error2")
	fmt.Printf("%T = %v\n", e1, e1) // *errors.errorString = this is error1
	fmt.Printf("%T = %v\n", e2, e2) // *errors.errorString = this is error2

	ea := errors.New("this is error")
	eb := errors.New("this is error")
	fmt.Println(ea == eb)                 // false
	fmt.Println(ea.Error() == eb.Error()) // true
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
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
