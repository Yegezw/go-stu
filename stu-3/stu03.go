package main

import (
	"fmt"
)

func test5() {
	fmt.Println("----------------------test5----------------------")

	arr := make([]int, 5, 10)
	fmt.Println(arr)

	// range
	for i, v := range arr {
		fmt.Printf("index = %d, value = %d\n", i, v)
	}

	// range _
	for _, v := range arr {
		fmt.Printf("value = %d\n", v)
	}
}
