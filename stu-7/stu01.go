package main

import (
	"fmt"
	"reflect"
)

func test1() {
	fmt.Println("----------------------test1----------------------")

	n := 1
	info(n)
	info(&n)
	fmt.Println()

	stu := Person{"zs", 18}
	info(stu)
	info(&stu)
	fmt.Println()

	var w WrapInt = 10
	info(w)
	info(&w)
}

type WrapInt int

type Person struct {
	Name string
	Age  int
}

func info(e any) {
	// 空接口 any = interface{} = (type, value)
	fmt.Printf("%v: %v, Kind = %v\n", reflect.TypeOf(e), reflect.ValueOf(e), reflect.TypeOf(e).Kind())
}
