package main

import (
	"fmt"
	"math"
)

func test4() {
	fmt.Println("----------------------test4----------------------")

	var i I

	i = F(math.Pi)
	i.M()
	describe(i) // (3.141592653589793, main.F)

	i = &T{S: "apple"}
	i.M()
	describe(i) // (&{apple}, *main.T)

	i = new(T)
	describe(i)         // (&{}, *main.T)
	i.(*T).S = "banana" // 类型断言
	i.M()
}

// ----------------------------------------------------------

// 接口值
// 接口也是值, 它们可以像其它值一样传递
// 接口值可以用作函数的参数或返回值

// 在内部, 接口值可以看做包含值和具体类型的元组 (value, type)
// 接口值保存了一个具体底层类型的具体值
// 接口值调用方法时会执行其底层类型的同名方法

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// ----------------------------------------------------------

type I interface {
	M()
}

// ----------------------------------------------------------

type F float64

// 值接收者
// 类型 F 实现了 I 接口

func (f F) M() {
	fmt.Println(f)
}

// ----------------------------------------------------------

type T struct {
	S string
}

// 指针接收者
// 类型 *T 实现了 I 接口

func (t *T) M() {
	fmt.Println(t.S)
}
