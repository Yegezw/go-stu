package main

import (
	"fmt"
)

func test3() {
	fmt.Println("----------------------test3----------------------")

	var animal Animal
	cat := Cat{
		Name: "cat",
		Age:  1,
	}
	dog := Dog{
		Name: "dog",
		Age:  1,
	}

	animal = cat
	animal = &cat // 也可以
	animal.AnimalInfo()

	// animal = dog 编译报错
	// func (dog *Dog) AnimalInfo() 为指针接收者
	// *Dog 实现了 Animal, Dog 未实现 Animal
	animal = &dog
	animal.AnimalInfo()
}

// ----------------------------------------------------------

// 接口类型: 一组方法签名定义的集合
// 接口类型的变量: 可以保存任何实现了这些方法的值

type Animal interface {
	AnimalInfo()
}

// ----------------------------------------------------------

type Cat struct {
	Name string
	Age  int
}

// 值接收者
// 类型 Cat 实现了 Animal 接口

func (cat Cat) AnimalInfo() {
	fmt.Printf("Name = %s, Age = %d\n", cat.Name, cat.Age)
}

// ----------------------------------------------------------

type Dog struct {
	Name string
	Age  int
}

// 指针接收者
// 类型 *Dog 实现了 Animal 接口

func (dog *Dog) AnimalInfo() {
	fmt.Printf("Name = %s, Age = %d\n", dog.Name, dog.Age)
}
