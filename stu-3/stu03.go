package main

import (
	"fmt"
	"sort"
)

func test6() {
	fmt.Println("----------------------test6----------------------")

	arr := make([]int, 5, 10)
	fmt.Println(arr)

	// range (v 不是 arr[i])
	for i, v := range arr {
		fmt.Printf("index = %d, value = %d\n", i, v)
	}

	// range _
	for _, v := range arr {
		fmt.Printf("value = %d\n", v)
	}

	/**
	 * [0] for range 会创建元素的副本, 直接修改迭代变量不会影响原切片
	 * [1] Go 1.22 之前通过 for _, v := range arr 遍历切片取不到变量的地址, 而是同一个临时变量的地址
	 * [2] for range 边遍历边 append 仍会停止
	 * [3] 闭包 {} 捕获引用, () 则捕获值
	 * [4] 遍历字典时, 顺序是随机的, 每次运行结果可能不同
	 * [5] for range 遍历字符串会返回 Unicode 代码点, 而不是字节
	 * [6] for range 中删除切片元素, 可能导致意外行为或漏掉某些元素
	 */

	// [1] 获取 arr1 内部元素地址, 给到 arr2 和 arr3
	arr1 := []int{1, 2}
	fmt.Println(&arr1[0], &arr1[1])

	var arr2 []*int
	for _, v := range arr1 {
		// 每次 v 都是同一个变量
		arr2 = append(arr2, &v)
	}
	fmt.Println(*arr2[0], *arr2[1]) // 2 2

	var arr3 []*int
	for i := range arr1 {
		arr3 = append(arr3, &arr1[i])
	}
	fmt.Println(*arr3[0], *arr3[1]) // 1 2

	// [2] for 循环仍会被停止
	v := []int{1, 2, 3}
	for _, value := range v {
		v = append(v, value)
	}
	fmt.Println(v)

	// [3] 闭包 {} 会捕获迭代变量的引用, 而不是它的值, () 则捕获值
	var funcs []func()
	for i := 1; i <= 3; i++ {
		i := i // 创建新的局部变量 i
		funcs = append(funcs, func() {
			fmt.Println(i)
		})
	}
	for _, f := range funcs {
		f() // 4 4 4
	}

	// [4] map 每次便利结果可能是不同的, 先对 key 排序后遍历
	dic := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := make([]string, 0, len(dic))
	for k := range dic {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("key = %s, value = %d\n", k, dic[k])
	}

	// [5] for range 遍历字符串会返回 Unicode 代码点, 而不是字节
	str := "hello 世界"
	for _, v := range str {
		fmt.Printf("char = %c\n", v)
	}
	for i := 0; i < len(str); i++ {
		fmt.Printf("index = %d, byte = %x\n", i, str[i])
	}
}

// 闭包 {} 捕获引用, () 则捕获值
func deferReturn1() {
	num := 1
	defer fmt.Println(num) // 1

	num = 2
	return
}

// 闭包 {} 捕获引用, () 则捕获值
func deferReturn2() {
	num := 1
	defer func() {
		fmt.Println(num) // 2
	}()

	num = 2
	return
}

// 设置返回值
// 执行 defer
// 将结果返回
func deferReturn3() (res int) {
	num := 1

	defer func() {
		res++
	}()

	return num // 2
}

// 设置返回值
// 执行 defer
// 将结果返回
func deferReturn4() (res int) {
	num := 1

	defer func() {
		num++
	}()

	return num // 1
}

// 设置返回值
// 执行 defer
// 将结果返回
func deferReturn5() int {
	num := 1

	defer func() {
		num++
	}()

	return num // 1
}
