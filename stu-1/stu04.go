package main

import (
	"fmt"
	"unsafe"
)

// 运算符
func test6() {
	fmt.Println("----------------------test6----------------------")

	num := 100

	var ptr *int
	ptr = &num

	fmt.Println("ptr 的大小为:", unsafe.Sizeof(ptr))
	fmt.Println("ptr 的值为:", ptr)
	fmt.Println("*ptr 的值为:", *ptr)

	*ptr = 1000
	fmt.Println("ptr 的值为:", ptr)
	fmt.Println("*ptr 的值为:", *ptr)

	// a &^ b = a & (^b)
	// a &^ b = 清除 a 中 ab 都为 1 的位
	a := 13     // 1101     a = 1101
	b := 11     // 1011    ^b = 0100
	c := a &^ b // 0100     c = 0100
	fmt.Println(c)
}
