package main

import "fmt"

func compute(fn func(int, int) int) int {
	return fn(1, 2)
}

func test6() {
	fmt.Println("----------------------test6----------------------")

	// 函数值
	sum := func(a int, b int) int {
		return a + b
	}
	fmt.Println(compute(sum))
}
