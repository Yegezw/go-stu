package main

import "fmt"

// for if-else switch-case
// defer + 释放资源 + 异常捕获
// range + deferReturn
// 函数变量 + 函数闭包
func main() {
	test1()
	test2()

	test3()
	test4()
	test5()

	test6()
	deferReturn1()
	deferReturn2()
	fmt.Println(deferReturn3())
	fmt.Println(deferReturn4())
	fmt.Println(deferReturn5())

	test7()
	test8()
}
