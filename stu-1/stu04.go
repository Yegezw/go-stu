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
}
