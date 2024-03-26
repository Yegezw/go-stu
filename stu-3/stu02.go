package main

import (
	"fmt"
)

func test3() {
	fmt.Println("----------------------test3----------------------")

	// defer 语句会将函数推迟到外层函数返回之后执行
	// 推迟调用的函数其参数会立即求值
	defer fmt.Println("world")

	fmt.Println("hello")
}

func test4() {
	fmt.Println("----------------------test4----------------------")

	fmt.Println("begin")

	// 推迟的函数调用会被压入一个栈中
	// 当外层函数返回时, 被推迟的函数会按照后进先出的顺序调用
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("end")
}
