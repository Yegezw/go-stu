package main

import "fmt"

// 常量计数器 iota, 被编译器修改的常量
func test5() {
	fmt.Println("----------------------test5----------------------")

	fmt.Println(a, b, c, d, e, f)
	fmt.Println(i, j, k, l)
}

// const 中每新增一行常量声明将使 iota 计数一次
const (
	a = iota
	b
	c
	d = 888  // iota++
	e = 999  // iota++
	f = iota // 恢复计数
)

// iota 在 const 关键字出现时将被重置为 O
const (
	i = 1 << iota // iota = 0
	j = 3 << iota // iota = 1
	k             // iota = 2
	l             // iota = 3
)
