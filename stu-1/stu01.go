package main

import (
	"fmt"
	"math"
)

// 变量
func test1() {
	fmt.Println("----------------------test1----------------------")

	/*
		bool

		string

		int  int8  int16  int32  int64
		uint uint8 uint16 uint32 uint64 uintptr

		byte // uint8 的别名

		rune // int32 的别名
			 // 表示一个 Unicode 码位

		float32 float64

		complex64 complex128

		类型转换
		i := 42
		f := float64(i)
		u := uint(f)

		%b 二进制
		%c 字符
		%s 字符串
		%t 布尔值
		%d 整数
		%g 小数
		%x 字节 (16 进制整数)
		%p 16 进制指针地址
		%T 类型
		%v 默认格式的值
		%q 带双引号的字符串
	*/

	var a int     // 默认 0      常用
	var b int = 1 // 初始化
	var c = 2     // 省略 type
	d := 3        // 省略 var    常用
	var e error
	fmt.Println(a, b, c, d, e)

	// 多变量声明
	num1, num2 := 1, 2
	fmt.Println(num1, num2)

	// 各 int 类型取值范围
	fmt.Println("uint8", 0, "~", math.MaxUint8)
	fmt.Println("int8 ", math.MinInt8, "~", math.MaxInt8)
	fmt.Println("int16", math.MinInt16, "~", math.MaxInt16)
	fmt.Println("int32", math.MinInt32, "~", math.MaxInt32)
	fmt.Println("int64", math.MinInt64, "~", math.MaxInt64)
}

func test2() {
	fmt.Println("----------------------test2----------------------")

	a := 1
	b := 2
	a, b = b, a // 交换
	fmt.Println(a, b)

	_, a = 9, 9 // _ 将会被丢弃, 它是只写的, 不能读
	fmt.Println(a)
}

// 全局变量必须用 var
// x, y := 1, 2 这种省略 var 的声明只能在函数中出现
// 生命周期: 全局变量是程序存活时间, 局部变量是函数存活时间
var (
	g1 int
	g2 bool
)
var g3, g4 int
var g5, g6 = "g5", "g6"

// 全局变量声明后可以不使用
// 函数中的变量声明后必须使用, 但函数中的常量声明后可以不用
