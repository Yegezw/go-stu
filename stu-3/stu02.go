package main

import (
	"fmt"
	"io"
	"os"
)

func test3() {
	fmt.Println("----------------------test3----------------------")

	// defer 语句会将函数推迟到外层函数返回之后执行 (return + panic)
	// 推迟调用的函数其参数会立即求值
	defer fmt.Println("world")

	fmt.Println("hello")
}

func test4() {
	fmt.Println("----------------------test4----------------------")

	fmt.Println("begin")

	// 推迟的函数调用会被压入一个栈中
	// 当外层函数返回时, 被推迟的函数会按照后进先出的顺序调用
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("end")
}

// 拷贝文件: defer 用来释放资源
func copyFile(src, dst string) (written int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer func() {
		err := srcFile.Close()
		if err != nil {
			return
		}
	}()

	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		err := dstFile.Close()
		if err != nil {
			return
		}
	}()

	return io.Copy(dstFile, srcFile)
}

// defer 配合 recover 一起处理 panic
func test5() {
	fmt.Println("----------------------test5----------------------")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	panic("hello panic")
}
