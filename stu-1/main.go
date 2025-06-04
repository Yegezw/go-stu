package main

import "fmt"

// 变量
// 常量
// 运算符
func main() {
	test1()
	test2()

	test3()
	test4()
	test5()

	test6()
}

// init 所在包首次加载时执行
func init() {
	fmt.Println("init")
}
