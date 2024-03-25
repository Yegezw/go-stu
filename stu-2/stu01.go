package main

import "fmt"

type People struct {
	Name string
	Age  int
}

type Student struct {
	ID    int
	Score float64
	People
}

// 结构体
func test1() {
	fmt.Println("----------------------test1----------------------")

	stu1 := Student{
		ID:    12,
		Score: 95.5,
		People: People{
			Name: "张三",
			Age:  18,
		},
	}
	fmt.Println("stu1:", stu1)
	fmt.Printf("stu1 的姓名是: %s\n", stu1.Name)

	// 其它属性为默认值
	stu2 := Student{
		People: People{Name: "李四"},
	}
	fmt.Printf("stu2: %v\n", stu2)
	fmt.Printf("stu2 的姓名是: %s\n", stu2.Name)

	stu3 := &Student{}
	fmt.Printf("stu3 的姓名是: %s\n", stu3.Name)
}
