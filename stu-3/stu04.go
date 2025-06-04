package main

import "fmt"

func compute(fn func(int, int) int) int {
	return fn(1, 2)
}

func test7() {
	fmt.Println("----------------------test7----------------------")

	// 函数变量
	sum := func(a int, b int) int {
		return a + b
	}
	fmt.Println(compute(sum))
}
