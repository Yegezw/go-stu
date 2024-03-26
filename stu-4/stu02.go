package main

import (
	"fmt"
)

func test2() {
	fmt.Println("----------------------test2----------------------")

	s := Student{
		Name: "小明",
		Age:  18,
	}
	s.Info()

	s.change1() // s.Age 未被改变
	s.Info()

	p := &s
	p.change2() // s.Age 被改变了
	s.Info()

	s.change2() // Go 会将它解释为 (&s).change2()
	p.Info()    // Go 会将它解释为 (*p).Info()

	// change3(s) 带指针参数的函数必须接受一个指针
}

// ----------------------------------------------------------

type Student struct {
	Name string
	Age  int
}

func (s Student) Info() {
	fmt.Printf("Name = %s, Age = %d\n", s.Name, s.Age)
}

func (s Student) change1() {
	s.Age++
}

// 指针接收者
// 以指针为接收者的方法被调用时, 接收者既能为值又能为指针
func (s *Student) change2() {
	s.Age++
}

// 带指针参数的函数必须接受一个指针
func change3(s *Student) {
	s.Age++
}

// ----------------------------------------------------------

// 使用指针接收者的原因有二
// 首先: 方法能够修改其接收者指向的值
// 其次: 这样可以避免在每次调用方法时复制该值, 若值的类型为大型结构体时, 这样做会更加高效
// 方法的接收者: 可以为值, 也可以为指针, 但并不应该二者混用
