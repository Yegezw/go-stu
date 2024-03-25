package main

import (
	"fmt"
	"unsafe"
)

// 常量
func test3() {
	fmt.Println("----------------------test3----------------------")

	const s1 string = "apple"
	const s2 = "banana" // 省略 type
	fmt.Println(s1, s2)

	// 多常量声明
	const a, b = 1, 2
	fmt.Println(a, b)
}

// 常量枚举 + 常量表达式
func test4() {
	fmt.Println("----------------------test4----------------------")

	fmt.Println(Unknown, Success, Fail)
	fmt.Println(p, q, r)
	fmt.Println(x1, x2)
}

// 常量枚举
const (
	Unknown = 1
	Success = 2
	Fail    = 3
)

// 常量表达式
// 常量可以用 len()、cap()、unsafe.Sizeof() 函数计算表达式的值
// 常量表达式中, 函数必须是内置函数, 否则编译不过
const (
	p = "apple"
	q = len(p)
	r = unsafe.Sizeof(p) // 16 byte = point(8) + len(8)
)

// x2 没有赋值, 那么 x2 将会是 x1 表达式的值
const (
	x1 = 9
	x2
)
