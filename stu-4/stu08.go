package main

import "fmt"

func test9() {
	fmt.Println("----------------------test9----------------------")

	a := Person{"王五", 33}
	z := Person{"赵六", 46}
	fmt.Println(a)
	fmt.Println(z)
}

// ----------------------------------------------------------

type Person struct {
	Name string
	Age  int
}

// fmt 包中定义的 Stringer 是最普遍的接口
// Stringer 是一个可以用字符串描述自己的类型

func (p Person) String() string {
	return fmt.Sprintf("Name = %s, Age = %d", p.Name, p.Age)
}
