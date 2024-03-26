package main

import (
	"fmt"
	"runtime"
)

func test1() {
	fmt.Println("----------------------test1----------------------")

	// for
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// while
	i := 1
	for i < 10 {
		i++
	}
	fmt.Println(i)

	// 无限循环
	// for {
	// 	fmt.Println(1)
	// }

	// if
	b := true
	if b {
		b = false
	}
	fmt.Println(b)

	// if
	if v := 101; !b {
		fmt.Println(v)
	}
}

func test2() {
	fmt.Println("----------------------test2----------------------")

	// switch case
	// 命中 case 后, 不会执行继续下面的 case
	// 除非 fallthrough
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("hello linux")
	case "windows":
		fmt.Println("hello windows")
		// fallthrough
	case "mac os":
		fmt.Println("max os")
		// fallthrough
	default:
		fmt.Println("I don't know!")
	}
}
