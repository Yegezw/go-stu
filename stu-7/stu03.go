package main

import (
	"fmt"
	"reflect"
)

// 可寻址就是可以找到数据本身, 而不是找到数据的副本
// 可寻址还不够, struct 里的未导出字段是不可设置值的(可修改)
func test3() {
	fmt.Println("----------------------test3----------------------")

	stu := &Student{
		Name:  "张三",
		Age:   20,
		score: 88.5,
	}
	fmt.Printf("stu = %+v\n", stu)

	v1 := reflect.ValueOf(stu)
	fmt.Println(v1.CanAddr())                 // false
	fmt.Println(v1.Elem().CanAddr())          // true
	fmt.Println(v1.Elem().Field(0).CanAddr()) // true

	slice := []int{1, 2, 3}
	v2 := reflect.ValueOf(slice)
	fmt.Println(v2.Index(0).CanAddr()) // true

	arr := [3]int{1, 2, 3}
	v3 := reflect.ValueOf(arr)
	fmt.Println(v3.Index(0).CanAddr()) // false

	fmt.Println(v1.Elem().Field(0).CanSet()) // Name  true
	fmt.Println(v1.Elem().Field(2).CanSet()) // score false

	v1.Elem().Field(0).SetString("李四")
	fmt.Printf("stu = %+v\n", stu)

	t := reflect.TypeOf(stu).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field %d: Name = %s, Type = %s, Tag = %s\n", i, field.Name, field.Type, field.Tag)
	}
}
