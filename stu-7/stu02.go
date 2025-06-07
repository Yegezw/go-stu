package main

import (
	"fmt"
	"reflect"
)

func test2() {
	fmt.Println("----------------------test2----------------------")

	stu := &Student{
		Name:  "张三",
		Age:   20,
		score: 88.5,
	}
	fmt.Printf("stu = %+v\n", stu)

	t := reflect.TypeOf(stu)
	v := reflect.ValueOf(stu)

	m1, _ := t.MethodByName("SetName")
	argsV1 := make([]reflect.Value, 0)
	argsV1 = append(argsV1, v)
	argsV1 = append(argsV1, reflect.ValueOf("李四"))
	m1.Func.Call(argsV1)
	fmt.Printf("stu = %+v\n", stu)

	m2 := v.MethodByName("SetName")
	argsV2 := make([]reflect.Value, 0)
	argsV2 = append(argsV2, reflect.ValueOf("王五"))
	m2.Call(argsV2)
	fmt.Printf("stu = %+v\n", stu)
}

// ----------------------------------------------------------

type Student struct {
	Name  string `json:"name"` // 标签 StructTag
	Age   int    `json:"age"`  // 标签 StructTag
	score float64
}

func (s *Student) GetName() string {
	return s.Name
}

func (s *Student) SetName(name string) {
	s.Name = name
}

func (s *Student) GetAge() int {
	return s.Age
}

func (s *Student) SetAge(age int) {
	s.Age = age
}

func (s *Student) GetScore() float64 {
	return s.score
}

func (s *Student) SetScore(score float64) {
	s.score = score
}
