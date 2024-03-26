package main

import (
	"fmt"
)

// 方法只是个带接收者参数的函数
// 接收者的类型定义和方法声明必须在同一包内, 不能为内建类型声明方法
func test1() {
	fmt.Println("----------------------test1----------------------")

	p := People{
		Name: "张三",
		Age:  21,
	}
	p.info()

	i := MyInt(10)
	i.info()
}

// ----------------------------------------------------------

type People struct {
	Name string
	Age  int
}

// 接收者为 People 的方法
func (p People) info() {
	fmt.Printf("Name = %s, Age = %d\n", p.Name, p.Age)
}

// ----------------------------------------------------------

type MyInt int64

// 接收者为 MyInt 的方法
// 也可以为非结构体类型声明方法
func (i MyInt) info() {
	fmt.Printf("MyInt = %d\n", i)
}
