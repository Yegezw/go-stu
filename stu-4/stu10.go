package main

import (
	"fmt"
	"io"
	"strings"
)

func test11() {
	fmt.Println("----------------------test11----------------------")

	// *strings.Reader &{Hello, Reader! 0 -1} 实现了 io/Reader
	r := strings.NewReader("Hello, Reader!")

	// 切片
	// len = 8, cap = 8
	b := make([]byte, 8)

	// Read(b)
	// 用数据填充给定的字节切片
	// 返回填充的字节数和错误值
	// 在遇到数据流的结尾时, 它会返回一个 io.EOF 错误
	for {
		n, err := r.Read(b) // 每次读取 8 字节
		fmt.Printf("b = %v, n = %v, err = %v\n", b, n, err)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
