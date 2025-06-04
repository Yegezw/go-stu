package main

import "fmt"

// 函数闭包
// 函数 adder 返回一个闭包, 每个闭包都被绑定在其各自的 sum 变量上
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func test8() {
	fmt.Println("----------------------test8----------------------")

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
