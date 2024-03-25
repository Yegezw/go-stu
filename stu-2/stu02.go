package main

import "fmt"

// 数组、切片、字典
func test2() {
	fmt.Println("----------------------test2----------------------")

	// ------------ 数组初始化 ------------

	var strArr = [10]string{}
	strArr[0] = "aa"
	strArr[1] = "bb"
	strArr[2] = "cc"
	strArr[3] = "dd"
	strArr[4] = "ee"
	fmt.Println(strArr)

	nums1 := [10]int{1, 2, 3}
	var nums2 = [...]int{1, 2, 3} // 推断元素个数
	fmt.Println("nums1 =", nums1)
	fmt.Println("nums2 =", nums2)

	// ------------ 切片初始化 ------------

	var slice1 []int
	slice1 = append(slice1, 1, 2)
	fmt.Println("slice1 =", slice1)

	var slice2 = []int{3, 4}
	fmt.Println("slice2 =", slice2)

	// 切片类型、初始化长度、容量
	var slice3 = make([]string, 0, 10)
	slice3 = strArr[1:3]
	fmt.Println("slice3 =", slice3)

	// 默认容量 = 长度
	slice4 := make([]int, 2)
	slice4 = append(slice4, 5, 6)
	fmt.Println("slice4 =", slice4)

	// ------------ 字典初始化 ------------

	var dic1 = map[string]int{
		"apple":  1,
		"banana": 2,
	}
	fmt.Println("dic1 =", dic1)

	dic2 := make(map[int]string)
	dic2[9] = "apple"
	dic2[6] = "banana"
	fmt.Println("dic2 =", dic2)
}
