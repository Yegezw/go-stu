package main

import (
	"fmt"
)

func test7() {
	fmt.Println("----------------------test7----------------------")

	var i interface{} = "hello"

	// 类型断言: 提供了访问接口值底层具体值的方式
	// 该语句断言接口值 i 保存了具体类型 string, 并将其底层类型为 string 的值赋予变量 s
	s := i.(string)
	fmt.Println(s)

	// 若 i 并未保存 string 类型的值, 该语句就会触发一个恐慌
	// 为了 "判断" 一个接口值是否保存了一个特定的类型, 类型断言可返回两个值: 底层值 + 报告断言是否成功的布尔值
	s, ok := i.(string)
	fmt.Println(s, ok)

	// 若 i 保存了一个 float64, 那么 f 将会是其底层值, 而 ok 为 true
	// 否则 ok 将为 false, 而 f 将为 float64 类型的零值, 程序并不会产生恐慌
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// 报错 panic
	// f = i.(float64)
	// fmt.Println(f)
}

func test8() {
	fmt.Println("----------------------test8----------------------")

	do(21)
	do("hello")
	do(true)
}

// ----------------------------------------------------------

// 类型选择: 按顺序从几个类型断言中选择分支的结构
// 类型选择与一般的 switch 语句相似
// 不过类型选择中的 case 为类型(而非值), 它们针对给定接口值所存储的值的类型进行比较

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%s is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
