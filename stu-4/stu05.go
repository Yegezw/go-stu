package main

import (
	"fmt"
)

func test5() {
	fmt.Println("----------------------test5----------------------")

	var i II

	// 正常情况
	i = &SS{S: "banana"}
	i.N()
	info(i)

	// 底层值为 nil 的接口值
	// 即便接口内的具体值为 nil, 方法仍然会被 nil 接收者调用
	var s *SS
	i = s
	i.N()
	info(i)

	// nil 接口值
	// nil 接口值既不保存值也不保存具体类型
	// 为 nil 接口调用方法会产生运行时错误, 因为接口的元组内并未包含能够指明该调用哪个 "具体" 方法的类型
	i = nil
	info(i)
	// i.N() 报错
}

// ----------------------------------------------------------

type II interface {
	N()
}

func info(i II) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// ----------------------------------------------------------

type SS struct {
	S string
}

func (t *SS) N() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}
