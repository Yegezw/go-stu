package main

import (
	"fmt"
)

func test6() {
	fmt.Println("----------------------test6----------------------")

	var i interface{}
	printInfo(i)

	i = 42
	printInfo(i)

	i = "hello world"
	printInfo(i)
}

// ----------------------------------------------------------

// 空接口被用来处理未知类型的值
// 指定了零个方法的接口值被称为 "空接口"
// 空接口可保存任何类型的值, 因为每个类型都至少实现了零个方法

func printInfo(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
